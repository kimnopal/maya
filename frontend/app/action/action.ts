import { posts, users } from "@/lib/DummyData";

export function findUser({ email }: { email: string }) {
  return users.find((user) => user.email == email);
}

export function getAllPosts({
  category,
  query,
}: {
  category: "competition" | "project" | undefined;
  query: string | undefined;
}) {
  return posts
    .filter((post) => post.category.includes(category ? category : ""))
    .filter((post) => post.title.includes(query ? query : ""));
}

export function findUserById(id: string) {
  return users.find((user) => user.id == id);
}
