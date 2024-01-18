"use client";

import * as z from "zod";
import React from "react";
import { Form, FormControl, FormField, FormItem, FormLabel } from "../ui/form";
import { LoginSchema } from "@/schemas";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Input } from "../ui/input";
import { Button } from "../ui/button";

const SigninForm = () => {
  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  const onSubmit = (values: z.infer<typeof LoginSchema>) => {
    console.log(values);
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
