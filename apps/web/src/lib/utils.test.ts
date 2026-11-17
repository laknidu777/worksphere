import { describe, it, expect } from "vitest";
import { formatLKR, formatSriLankaDate } from "./utils";

describe("Sri Lankan Business Utilities", () => {
  it("formats LKR currency correctly with 2 decimal places", () => {
    const formatted = formatLKR(150000.5);
    // Should include LKR and 150,000.50
    expect(formatted).toContain("150,000.50");
  });

  it("formats Sri Lankan date according to DD/MM/YYYY convention", () => {
    // 2026-11-17T00:00:00Z -> in Asia/Colombo is 17/11/2026
    const date = new Date("2026-11-17T06:00:00Z");
    const formatted = formatSriLankaDate(date);
    expect(formatted).toBe("17/11/2026");
  });
});
