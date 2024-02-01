// "use client";

import { getAuth } from "@/lib/cookie";
import Image from "next/image";
import React from "react";

const UserProfile = async () => {
  const authUser = await getAuth();
  console.log(authUser);

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
      <h1 className="text-xl font-bold text">{authUser?.name}</h1>
      <p>{authUser?.email}</p>
      <p>{authUser?.majors?.name}</p>
      <p>{authUser?.majors?.faculty}</p>
      <button
        // onClick={() => {}}
        className="font-bold py-2 bg-secondary rounded-xl"
      >
        Sign out
      </button>
    </div>
  );
};

export default UserProfile;
