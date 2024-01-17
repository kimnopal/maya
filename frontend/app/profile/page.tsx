import { getServerSession } from "next-auth";
import React from "react";
import { options } from "../api/auth/[...nextauth]/options";
import { redirect } from "next/navigation";
import BackTemplate from "@/components/template/BackTemplate";
import UserProfile from "@/components/cards/UserProfile";

const ProfilePage = async () => {
  const session = await getServerSession(options);

  if (!session) {
    redirect("/api/auth/signin?callbackUrl=/profile");
  }

  return (
    <BackTemplate>
      <UserProfile />
    </BackTemplate>
  );
};

export default ProfilePage;
