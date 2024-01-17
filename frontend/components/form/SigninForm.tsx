"use client";
import { signIn } from "next-auth/react";
import React, { useRef } from "react";
import TextBox from "./TextBox";

const SigninForm = async () => {
  const userEmail = useRef("");
  const pass = useRef("");

  const onSubmit = async () => {
    const result = await signIn("credentials", {
      email: userEmail.current,
      password: pass.current,
      redirect: true,
      callbackUrl: "/",
    });
  };

  return (
    <>
      <TextBox
        labelText="Email Address"
        type="email"
        onChange={(e) => {
          userEmail.current = e.target.value;
        }}
        placeholder="youlooknicetoday@example.com"
      />
      <TextBox
        labelText="Password"
        type="password"
        onChange={(e) => {
          pass.current = e.target.value;
        }}
        placeholder="Ssst!"
      />
      <button
        onClick={onSubmit}
        className="font-bold text-white bg-primary py-3 rounded-xl"
      >
        Sign in
      </button>
    </>
  );
};

export default SigninForm;
