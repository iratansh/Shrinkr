import { useState } from "react";
import { ShortenForm } from "./components/ShortenForm";
import { ResultCard } from "./components/ResultCard";
import { AnalyticsTable } from "./components/AnalyticsTable";

export function App() {
  const [lastShortUrl, setLastShortUrl] = useState<string | null>(null);

  return (
    <main style={{ maxWidth: 720, margin: "2rem auto", fontFamily: "system-ui" }}>
      <h1>Shrinkr</h1>
      <ShortenForm onShortened={setLastShortUrl} />
      {lastShortUrl && <ResultCard shortUrl={lastShortUrl} />}
      <AnalyticsTable />
    </main>
  );
}
