import { findUser } from "@/app/action/action";
import { userType } from "@/app/types/type";
import { NextResponse } from "next/server";

export async function POST(req: Request) {
  const { email } = await req.json();

  const foundUser: userType | undefined = findUser({ email });

  return NextResponse.json(foundUser, { status: 201 });
}
