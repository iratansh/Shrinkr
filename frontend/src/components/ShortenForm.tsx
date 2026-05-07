import { useState } from "react";
import { shorten } from "../api/client";

interface Props {
  onShortened: (shortUrl: string) => void;
}

export function ShortenForm({ onShortened }: Props) {
  const [url, setUrl] = useState("");
  const [busy, setBusy] = useState(false);
  const [error, setError] = useState<string | null>(null);

  async function onSubmit(e: React.FormEvent) {
    e.preventDefault();
    setBusy(true);
    setError(null);
    try {
      const res = await shorten(url);
      onShortened(res.short_url);
      setUrl("");
    } catch (err) {
      setError((err as Error).message);
    } finally {
      setBusy(false);
    }
  }

  return (
    <form onSubmit={onSubmit}>
      <input
        type="url"
        required
        placeholder="https://example.com/long/url"
        value={url}
        onChange={(e) => setUrl(e.target.value)}
        style={{ width: "100%", padding: "0.5rem" }}
      />
      <button type="submit" disabled={busy} style={{ marginTop: "0.5rem" }}>
        {busy ? "Shortening..." : "Shorten"}
      </button>
      {error && <p style={{ color: "crimson" }}>{error}</p>}
    </form>
  );
}
