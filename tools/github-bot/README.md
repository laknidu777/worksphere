# WorkSphere GitHub + OpenAI integration

This starter project adds a small GitHub webhook receiver that listens for issues and uses OpenAI to draft a response. It is designed as a foundation for WorkSphere's GitHub automation.

## Features

- Express webhook endpoint for GitHub events
- Signature validation using `GITHUB_WEBHOOK_SECRET`
- OpenAI-powered issue response generation
- GitHub App installation authentication to post comments back to the repository
- Health check at `/health`

## Prerequisites

- Node.js 18+
- A GitHub App configured with:
  - App ID
  - Private key
  - Webhook secret
  - Repository installation
- An OpenAI API key

## Installation

```bash
npm install
```

## Environment variables

Copy `.env.example` to `.env` and update the values:

```bash
cp .env.example .env
```

## Start the app

```bash
npm start
```

## GitHub webhook setup

1. Create a GitHub App.
2. Set the webhook URL to your deployed app endpoint, for example:
   `https://your-domain.example.com/api/webhooks/github`
3. Enable the `issues` event.
4. Install the app on the target repository.
5. Add the app's App ID, private key, and webhook secret to your environment.

## Notes

This is intentionally a lightweight starter integration rather than a full production system. It is a strong base for adding PR summaries, feature requests, triage, and AI assistance workflows.
