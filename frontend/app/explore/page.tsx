import CateogryBar from "@/components/navigation/CategoryBar";
import BackTemplate from "@/components/template/BackTemplate";
import React from "react";
import { getAllPosts } from "../action/action";
import ItemList from "@/components/cards/ItemList";

const ExplorePage = ({
  searchParams,
}: {
  searchParams: {
    query?: string;
    category?: "competition" | "project" | undefined;
  };
}) => {
  const posts = getAllPosts({
    category: searchParams?.category,
    query: searchParams?.query,
  });

  return (
    <BackTemplate>
      <CateogryBar />
      <ul className="flex flex-col gap-2">
        {posts.map((post) => (
          <ItemList item={post} key={post.id} />
        ))}
      </ul>
    </BackTemplate>
  );
};

export default ExplorePage;
