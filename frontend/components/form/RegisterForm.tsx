"use client";
import React, { useRef } from "react";
import TextBox from "./TextBox";

const RegisterForm = async () => {
  const userEmail = useRef("");
  const pass = useRef("");

  const onSubmit = async () => {};
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
        Sign up
      </button>
    </>
  );
};

export default RegisterForm;
