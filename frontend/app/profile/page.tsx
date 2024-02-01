import React from "react";
import { options } from "../api/auth/nextauth/options";
import { redirect } from "next/navigation";
import BackTemplate from "@/components/template/BackTemplate";
import UserProfile from "@/components/cards/UserProfile";

const ProfilePage = async () => {
  return (
    <BackTemplate>
      <UserProfile />
    </BackTemplate>
  );
};

export default ProfilePage;
