import Link from "next/link";

export default function NotFound() {
  return (
    <div className="h-full min-h-screen w-full flex flex-col gap-3 justify-center items-center text-center p-4">
      <h1 className="font-bold text-xl">Not found – 404!</h1>
      <p>
        I'm so sorry there must have been a mistake or we are currently still
        working this feature. <br /> Thank you for your understanding! ❤️
      </p>
      <Link
        href="/"
        className="px-4 py-2 font-bold text-white bg-primary text-lg rounded-md"
      >
        Back to Home
      </Link>
    </div>
  );
}
