const API_BASE = import.meta.env.VITE_API_BASE ?? "/api";

export interface ShortenResponse {
  code: string;
  short_url: string;
}

export interface AnalyticsResponse {
  code: string;
  total_clicks: number;
}

export async function shorten(url: string): Promise<ShortenResponse> {
  const res = await fetch(`${API_BASE}/shorten`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ url }),
  });
  if (!res.ok) throw new Error(`shorten failed: ${res.status}`);
  return res.json();
}

export async function getAnalytics(code: string): Promise<AnalyticsResponse> {
  const res = await fetch(`${API_BASE}/analytics/${encodeURIComponent(code)}`);
  if (!res.ok) throw new Error(`analytics failed: ${res.status}`);
  return res.json();
}
