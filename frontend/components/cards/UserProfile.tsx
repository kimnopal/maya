"use client";
import { signOut, useSession } from "next-auth/react";
import Image from "next/image";
import React from "react";

const UserProfile = () => {
  const { data: session } = useSession({
    required: true,
  });

  console.log(session);

  return (
    <div className="flex flex-col gap-2 px-3 py-6 text-white bg-secondary-foreground rounded-2xl">
      <div className="flex justify-center items-center">
        <Image
          src="/saujana.jpg"
          alt="profile"
          width={128}
          height={128}
          className="rounded-full"
        />
      </div>
      <h1 className="text-xl font-bold text">{session?.user?.name}</h1>
      <p>{session?.user?.email}</p>
      <p>{session?.user?.majors.name}</p>
      <p>{session?.user?.majors.faculty}</p>
      <button
        onClick={() => {
          signOut();
        }}
        className="font-bold py-2 bg-secondary rounded-xl"
      >
        Sign out
      </button>
    </div>
  );
};

export default UserProfile;
