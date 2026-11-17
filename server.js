const express = require('express');
const crypto = require('crypto');
const { Octokit } = require('@octokit/rest');
const { createAppAuth } = require('@octokit/auth-app');
const OpenAI = require('openai');
require('dotenv').config();

const app = express();
const port = process.env.PORT || 3000;

app.use(express.raw({ type: 'application/json' }));

function verifyGithubSignature(rawBody, signature) {
  const secret = process.env.GITHUB_WEBHOOK_SECRET;

  if (!secret) {
    return true;
  }

  if (!signature) {
    return false;
  }

  const expected = `sha256=${crypto
    .createHmac('sha256', secret)
    .update(rawBody)
    .digest('hex')}`;

  return crypto.timingSafeEqual(
    Buffer.from(expected),
    Buffer.from(signature)
  );
}

function getOpenAIClient() {
  if (!process.env.OPENAI_API_KEY) {
    throw new Error('OPENAI_API_KEY is not set.');
  }

  return new OpenAI({ apiKey: process.env.OPENAI_API_KEY });
}

async function generateIssueReply(issue) {
  const client = getOpenAIClient();
  const issueText = [
    issue.title,
    issue.body || 'No description provided.'
  ].join('\n\n');

  const completion = await client.chat.completions.create({
    model: 'gpt-4o-mini',
    messages: [
      {
        role: 'system',
        content:
          'You are a helpful support and engineering assistant for WorkSphere, a Sri Lankan multi-tenant SaaS platform. Respond briefly, accurately, and in a professional tone. If the issue is vague, ask one clarifying follow-up question and suggest the next best step.'
      },
      {
        role: 'user',
        content: `Draft a helpful GitHub issue response for this WorkSphere report:\n\n${issueText}`
      }
    ],
    temperature: 0.7,
    max_tokens: 300
  });

  return completion.choices[0]?.message?.content?.trim() || 'Thanks for raising this. We are reviewing it.';
}

async function createInstallationOctokit(installationId) {
  const appId = process.env.GITHUB_APP_ID;
  const privateKey = process.env.GITHUB_PRIVATE_KEY;

  if (!appId || !privateKey) {
    throw new Error('Github App credentials are not configured.');
  }

  return new Octokit({
    authStrategy: createAppAuth,
    auth: {
      appId: Number(appId),
      privateKey: privateKey.replace(/\\n/g, '\n'),
      installationId: Number(installationId)
    }
  });
}

async function addCommentToIssue(payload, body) {
  const installationId = payload.installation?.id;

  if (!installationId) {
    throw new Error('GitHub installation context missing.');
  }

  const octokit = await createInstallationOctokit(installationId);
  const { owner, repo, number } = payload.repository;

  await octokit.rest.issues.createComment({
    owner: owner.login,
    repo: repo.name,
    issue_number: number,
    body
  });
}

app.get('/health', (req, res) => {
  res.json({ ok: true, message: 'WorkSphere GitHub + OpenAI integration is running.' });
});

app.post('/api/webhooks/github', async (req, res) => {
  const rawBody = req.body;
  const signature = req.headers['x-hub-signature-256'];

  if (typeof rawBody === 'undefined' || rawBody.length === 0) {
    return res.status(400).json({ error: 'Empty request body.' });
  }

  if (!verifyGithubSignature(rawBody, signature)) {
    return res.status(401).json({ error: 'Invalid GitHub signature.' });
  }

  try {
    const eventName = req.headers['x-github-event'];
    const payload = JSON.parse(rawBody.toString('utf8'));

    if (eventName === 'issues' && payload.action === 'opened') {
      const responseText = await generateIssueReply(payload.issue);
      await addCommentToIssue(payload, responseText);
    }

    return res.status(200).json({ ok: true, received: eventName });
  } catch (error) {
    console.error('GitHub webhook error:', error);
    return res.status(500).json({
      ok: false,
      error: error.message || 'Unexpected error while processing the webhook.'
    });
  }
});

app.listen(port, () => {
  console.log(`WorkSphere GitHub integration running on http://localhost:${port}`);
});
