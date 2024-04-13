export default function Home() {
  return (
    <main className="flex gap-3">
      HOLA MUNDO
      
      <video controls poster="favicon.ico" className="aspect-video" width={720}>
        <source src="https://storage.droque.dev/1-minute-sample-video.mp4" type="video/mp4" />
        Your browser does not support the video tag.
      </video>
    </main>
  );
}