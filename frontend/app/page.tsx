import ItemList from "@/components/ItemList";
import Navbar from "@/components/Navbar";
import SearchBar from "@/components/SearchBar";
import UserStats from "@/components/UserStats";
import { posts, user } from "@/lib/DummyData";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  return (
    <>
      <div className="flex flex-col gap-6 px-6 pt-6 py-24">
        <SearchBar />
        <main className="flex flex-col gap-3 text-black">
          <UserStats user={user} />

          <h2 className="font-bold text-xl">Explore</h2>

          <div className="grid grid-flow-col grid-cols-2 gap-3 text-white bg-white p-4 rounded-2xl">
            <Link
              href="/explore?category=project"
              className="flex flex-col justify-center items-center py-5 bg-primary rounded-2xl gap-3"
            >
              <Image width={64} height={64} src="/project.png" alt="Project" />
              <h4 className="text-xl font-semibold">Project</h4>
            </Link>
            <Link
              href="/explore?category=project"
              className="flex flex-col justify-center items-center py-5 bg-secondary rounded-2xl gap-3"
            >
              <Image
                width={64}
                height={64}
                src="/competition.png"
                alt="Project"
              />
              <h4 className="text-xl font-semibold">Competition</h4>
            </Link>
          </div>

          <h2 className="font-bold text-xl">Our latest update</h2>

          {posts.map((post, index) => {
            return <ItemList item={post} key={index} />;
          })}
        </main>
      </div>
      <Navbar />
    </>
  );
}
