import React from "react";
import {
  Building2,
  FolderKanban,
  Clock,
  Users2,
  DollarSign,
  BarChart3,
  Bell,
  CreditCard,
  Mail,
  ShieldCheck,
  Globe2,
  Server,
  Layers,
  ChevronRight,
  Sparkles
} from "lucide-react";

const modules = [
  {
    id: "orgs",
    name: "Organization & Users",
    description: "Multi-tenant workspace setup, departments, employee directories, and RBAC permissions.",
    icon: Building2,
    badge: "Foundation Ready",
    color: "from-blue-500/10 to-indigo-500/10 border-blue-500/20 text-blue-600 dark:text-blue-400"
  },
  {
    id: "projects",
    name: "Projects & Tasks",
    description: "Kanban boards, task lists, milestone tracking, priorities, deadlines, and dependencies.",
    icon: FolderKanban,
    badge: "Planned",
    color: "from-violet-500/10 to-purple-500/10 border-violet-500/20 text-violet-600 dark:text-violet-400"
  },
  {
    id: "workforce",
    name: "Workforce & Time",
    description: "Digital clock-in/out, timesheet approvals, overtime, leave entitlements, and workload tracking.",
    icon: Clock,
    badge: "Planned",
    color: "from-emerald-500/10 to-teal-500/10 border-emerald-500/20 text-emerald-600 dark:text-emerald-400"
  },
  {
    id: "crm",
    name: "Client & CRM",
    description: "Customer directories, lead pipelines, interaction history, and restricted client portals.",
    icon: Users2,
    badge: "Planned",
    color: "from-amber-500/10 to-orange-500/10 border-amber-500/20 text-amber-600 dark:text-amber-400"
  },
  {
    id: "finance",
    name: "Finance & Expenses",
    description: "LKR-denominated invoices, expense claims, project budgets, and real-time profitability margins.",
    icon: DollarSign,
    badge: "Planned",
    color: "from-green-500/10 to-emerald-500/10 border-green-500/20 text-green-600 dark:text-green-400"
  },
  {
    id: "analytics",
    name: "Business Analytics",
    description: "Executive revenue summaries, employee utilization, overdue invoice tracking, and business KPIs.",
    icon: BarChart3,
    badge: "Planned",
    color: "from-cyan-500/10 to-blue-500/10 border-cyan-500/20 text-cyan-600 dark:text-cyan-400"
  },
  {
    id: "notifications",
    name: "Notifications & Automation",
    description: "Transactional in-app alerts, email reminders, leave routing, and scheduled automation.",
    icon: Bell,
    badge: "Planned",
    color: "from-pink-500/10 to-rose-500/10 border-pink-500/20 text-pink-600 dark:text-pink-400"
  },
  {
    id: "billing",
    name: "SaaS Subscriptions",
    description: "Multi-tenant plans, trial enforcement, usage limits, and candidate PayHere IPG integration.",
    icon: CreditCard,
    badge: "Proposed / Candidate IPG",
    color: "from-indigo-500/10 to-sky-500/10 border-indigo-500/20 text-indigo-600 dark:text-indigo-400"
  },
  {
    id: "email",
    name: "Email Marketing",
    description: "Outbound campaign authoring, audience segments, deliverability telemetry, and suppression lists.",
    icon: Mail,
    badge: "Outbound Only (Candidate SES)",
    color: "from-sky-500/10 to-teal-500/10 border-sky-500/20 text-sky-600 dark:text-sky-400"
  }
];

export default function DashboardPage() {
  return (
    <div className="flex min-h-screen bg-slate-900 text-slate-100">
      {/* Sidebar */}
      <aside className="w-64 border-r border-slate-800 bg-slate-950/70 p-5 flex flex-col justify-between hidden md:flex">
        <div>
          <div className="flex items-center gap-3 px-2 py-3 mb-6">
            <div className="flex h-10 w-10 items-center justify-center rounded-xl bg-gradient-to-tr from-indigo-600 to-indigo-400 text-white font-bold shadow-lg shadow-indigo-500/25">
              WS
            </div>
            <div>
              <span className="font-bold text-lg tracking-tight text-white block">WorkSphere</span>
              <span className="text-xs text-indigo-400 font-medium tracking-wide uppercase">Operations SaaS</span>
            </div>
          </div>

          <nav className="space-y-1">
            <div className="px-3 py-2 text-xs font-semibold text-slate-400 uppercase tracking-wider">
              Core Modules
            </div>
            {modules.map((m) => {
              const Icon = m.icon;
              return (
                <div
                  key={m.id}
                  className="flex items-center gap-3 px-3 py-2 rounded-lg text-sm text-slate-300 hover:bg-slate-800/60 hover:text-white transition-colors cursor-pointer group"
                >
                  <Icon className="h-4 w-4 text-slate-400 group-hover:text-indigo-400 transition-colors" />
                  <span className="truncate">{m.name}</span>
                </div>
              );
            })}
          </nav>
        </div>

        <div className="rounded-xl border border-slate-800 bg-slate-900/50 p-3 text-xs space-y-2">
          <div className="flex items-center gap-2 text-slate-400">
            <Globe2 className="h-3.5 w-3.5 text-indigo-400" />
            <span>Sri Lanka Defaults</span>
          </div>
          <div className="text-[11px] text-slate-400 space-y-0.5 pl-5">
            <div>Currency: <span className="text-slate-200">LKR</span></div>
            <div>Timezone: <span className="text-slate-200">Asia/Colombo</span></div>
            <div>Locale: <span className="text-slate-200">en-LK (DD/MM/YYYY)</span></div>
          </div>
        </div>
      </aside>

      {/* Main Content Area */}
      <main className="flex-1 flex flex-col min-w-0">
        {/* Top Header */}
        <header className="h-16 border-b border-slate-800 bg-slate-950/40 px-6 flex items-center justify-between">
          <div className="flex items-center gap-3">
            <div className="md:hidden flex h-8 w-8 items-center justify-center rounded-lg bg-indigo-600 text-white font-bold text-sm">
              WS
            </div>
            <h1 className="text-base font-semibold text-white">System Architecture & Operations Dashboard</h1>
          </div>
          <div className="flex items-center gap-3">
            <span className="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-xs font-medium bg-emerald-500/10 text-emerald-400 border border-emerald-500/20">
              <span className="h-1.5 w-1.5 rounded-full bg-emerald-400 animate-pulse"></span>
              Modular Monolith Foundation
            </span>
          </div>
        </header>

        {/* Content Body */}
        <div className="p-6 md:p-8 space-y-8 max-w-7xl mx-auto w-full">
          {/* Welcome Banner */}
          <div className="relative overflow-hidden rounded-2xl border border-slate-800 bg-gradient-to-r from-slate-900 via-indigo-950/40 to-slate-900 p-6 md:p-8">
            <div className="relative z-10 max-w-3xl space-y-3">
              <div className="inline-flex items-center gap-2 rounded-full border border-indigo-500/30 bg-indigo-500/10 px-3 py-1 text-xs text-indigo-300">
                <Sparkles className="h-3.5 w-3.5 text-indigo-400" />
                WorkSphere Enterprise Platform Architecture
              </div>
              <h2 className="text-2xl md:text-3xl font-bold tracking-tight text-white">
                Sri Lanka-First Multi-Tenant Business Operations Platform
              </h2>
              <p className="text-slate-400 text-sm leading-relaxed">
                WorkSphere consolidates organizational management, workforce timekeeping, client relationships, project delivery, and financial operations into a high-performance modular monolith built on Go, Next.js, and PostgreSQL with Row-Level Security.
              </p>
            </div>
          </div>

          {/* Quick Metrics Bar (Sri Lanka Defaults) */}
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 gap-4">
            <div className="rounded-xl border border-slate-800 bg-slate-950/40 p-4 space-y-1">
              <div className="text-xs text-slate-400 flex items-center justify-between">
                <span>Architecture</span>
                <Layers className="h-4 w-4 text-indigo-400" />
              </div>
              <div className="text-lg font-bold text-white">Modular Monolith</div>
              <div className="text-[11px] text-slate-400">Go 1.23 + Chi + sqlc</div>
            </div>

            <div className="rounded-xl border border-slate-800 bg-slate-950/40 p-4 space-y-1">
              <div className="text-xs text-slate-400 flex items-center justify-between">
                <span>Multi-Tenancy</span>
                <ShieldCheck className="h-4 w-4 text-emerald-400" />
              </div>
              <div className="text-lg font-bold text-white">Shared Schema</div>
              <div className="text-[11px] text-slate-400">PostgreSQL RLS + organization_id</div>
            </div>

            <div className="rounded-xl border border-slate-800 bg-slate-950/40 p-4 space-y-1">
              <div className="text-xs text-slate-400 flex items-center justify-between">
                <span>Localization</span>
                <Globe2 className="h-4 w-4 text-amber-400" />
              </div>
              <div className="text-lg font-bold text-white">LKR / Asia-Colombo</div>
              <div className="text-[11px] text-slate-400">en-LK format • UTC Storage</div>
            </div>

            <div className="rounded-xl border border-slate-800 bg-slate-950/40 p-4 space-y-1">
              <div className="text-xs text-slate-400 flex items-center justify-between">
                <span>Background Tasks</span>
                <Server className="h-4 w-4 text-sky-400" />
              </div>
              <div className="text-lg font-bold text-white">River + Redis</div>
              <div className="text-[11px] text-slate-400">Transactional database queues</div>
            </div>
          </div>

          {/* Module Grid */}
          <div className="space-y-4">
            <div className="flex items-center justify-between">
              <h3 className="text-lg font-semibold text-white">Core Business Modules</h3>
              <span className="text-xs text-slate-400">9 Core Domains Defined</span>
            </div>

            <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
              {modules.map((mod) => {
                const Icon = mod.icon;
                return (
                  <div
                    key={mod.id}
                    className="group relative rounded-xl border border-slate-800 bg-slate-950/40 p-5 hover:border-slate-700 transition-all hover:shadow-lg hover:shadow-indigo-500/5 flex flex-col justify-between space-y-4"
                  >
                    <div className="space-y-3">
                      <div className="flex items-start justify-between">
                        <div className={`p-2.5 rounded-lg border bg-gradient-to-br ${mod.color}`}>
                          <Icon className="h-5 w-5" />
                        </div>
                        <span className="text-[11px] font-medium px-2 py-0.5 rounded-full border border-slate-800 bg-slate-900 text-slate-400">
                          {mod.badge}
                        </span>
                      </div>
                      <div>
                        <h4 className="font-semibold text-white group-hover:text-indigo-300 transition-colors">
                          {mod.name}
                        </h4>
                        <p className="text-xs text-slate-400 mt-1 line-clamp-2 leading-relaxed">
                          {mod.description}
                        </p>
                      </div>
                    </div>
                    <div className="pt-2 border-t border-slate-800/60 flex items-center justify-between text-xs text-slate-400">
                      <span className="font-mono text-[11px] text-slate-400">internal/{mod.id}</span>
                      <ChevronRight className="h-4 w-4 text-slate-400 group-hover:translate-x-0.5 transition-transform" />
                    </div>
                  </div>
                );
              })}
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
