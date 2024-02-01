import { cookies } from "next/headers";
import { NextRequest, NextResponse } from "next/server";

export async function POST(request: NextRequest) {
    const body = await request.json()
    const response = fetch("http://localhost:8000/api/users/login", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify({
            username: body.username,
            password: body.password,
        }),
    })

    const responseJson = await (await response).json()

    cookies().set("token", responseJson.data.token, {
        maxAge: 60 * 60 * 48,
    })

    return NextResponse.json(responseJson)
}