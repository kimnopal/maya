"use client";
import { useRouter } from "next/navigation";
import React from "react";

const SigninButton = () => {
  const router = useRouter();

  return (
    <button
      className="text-center font-bold py-3 rounded-xl bg-primary"
      onClick={() => router.push("/signin")}
    >
      Sign in
    </button>
  );
};

export default SigninButton;
