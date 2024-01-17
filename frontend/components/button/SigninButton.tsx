"use client";
import { signIn } from "next-auth/react";
import React from "react";

const SigninButton = () => {
  return (
    <button
      className="text-center font-bold py-3 rounded-xl bg-primary"
      onClick={() => signIn()}
    >
      Sign in
    </button>
  );
};

export default SigninButton;
