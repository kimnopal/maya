"use client";

import * as z from "zod";
import React from "react";
import { Form, FormControl, FormField, FormItem, FormLabel } from "../ui/form";
import { LoginSchema } from "@/schemas";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import { Input } from "../ui/input";
import { Button } from "../ui/button";
import axios from "axios";
import { useRouter } from "next/navigation";

const SigninForm = () => {
  const form = useForm<z.infer<typeof LoginSchema>>({
    resolver: zodResolver(LoginSchema),
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const router = useRouter();

  const onSubmit = async (values: z.infer<typeof LoginSchema>) => {
    try {
      const data = await axios.post(
        "http://localhost:8000/api/users/login",
        JSON.stringify(values),
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );

      router.push("/");
    } catch (error) {
      console.log(error);
    }
  };

  return (
    <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
        <FormField
          control={form.control}
          name="username"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-base">Username</FormLabel>
              <FormControl>
                <Input {...field} placeholder="youlooknicetoday" type="text" />
              </FormControl>
            </FormItem>
          )}
        />
        <FormField
          control={form.control}
          name="password"
          render={({ field }) => (
            <FormItem>
              <FormLabel className="text-base">Password</FormLabel>
              <FormControl>
                <Input {...field} placeholder="Ssst!" type="password" />
              </FormControl>
            </FormItem>
          )}
        />
        <Button type="submit" className="w-full font-bold rounded-xl">
          Sign in
        </Button>
      </form>
    </Form>
  );
};

export default SigninForm;
