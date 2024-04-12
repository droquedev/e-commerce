import Image from "next/image";

async function getData() {
  const res = await fetch("http://localhost:8000/books");
  const data = await res.json();
  return data.books as { title: string; author: string, foto: string }[];
}

export default async function Home() {
  const data = await getData();

  console.log(data)
  return (
    <main className="flex gap-3">
      {
        data.map((book) => (
          <div key={book.title} className="border border-white rounded-md p-2 max-w-[210px]">
            <h1>{book.title}</h1>
            <h2>{book.author}</h2>
            <img alt={`${book.title} ${book.author}`} src={book.foto} width={200} height={200}     />
          </div>
        ))
      }
    </main>
  );
}
