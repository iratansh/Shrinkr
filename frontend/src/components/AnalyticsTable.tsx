import { useState } from "react";
import { getAnalytics, AnalyticsResponse } from "../api/client";

export function AnalyticsTable() {
  const [code, setCode] = useState("");
  const [data, setData] = useState<AnalyticsResponse | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function onLookup() {
    setError(null);
    try {
      setData(await getAnalytics(code));
    } catch (err) {
      setError((err as Error).message);
    }
  }

  return (
    <section style={{ marginTop: "2rem" }}>
      <h2>Analytics</h2>
      <input
        placeholder="short code"
        value={code}
        onChange={(e) => setCode(e.target.value)}
      />
      <button onClick={onLookup} style={{ marginLeft: "0.5rem" }}>Look up</button>
      {error && <p style={{ color: "crimson" }}>{error}</p>}
      {data && (
        <table style={{ marginTop: "0.5rem", borderCollapse: "collapse" }}>
          <tbody>
            <tr>
              <th style={{ textAlign: "left", paddingRight: "1rem" }}>Code</th>
              <td>{data.code}</td>
            </tr>
            <tr>
              <th style={{ textAlign: "left", paddingRight: "1rem" }}>Total clicks</th>
              <td>{data.total_clicks}</td>
            </tr>
          </tbody>
        </table>
      )}
    </section>
  );
}
