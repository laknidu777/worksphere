import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "WorkSphere — Business Operations Platform",
  description: "Sri Lanka-first, subscription-based, multi-tenant business operations SaaS platform."
};

export default function RootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return (
    <html lang="en">
      <body className="min-h-screen bg-slate-50 text-slate-900 antialiased dark:bg-slate-950 dark:text-slate-50">
        {children}
      </body>
    </html>
  );
}
