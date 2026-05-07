interface Props {
  shortUrl: string;
}

export function ResultCard({ shortUrl }: Props) {
  return (
    <section style={{ marginTop: "1rem", padding: "1rem", border: "1px solid #ddd", borderRadius: 8 }}>
      <h2>Your short link</h2>
      <a href={shortUrl} target="_blank" rel="noreferrer">{shortUrl}</a>
      <button
        onClick={() => navigator.clipboard.writeText(shortUrl)}
        style={{ marginLeft: "0.5rem" }}
      >
        Copy
      </button>
    </section>
  );
}
