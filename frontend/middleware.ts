import { NextRequest, NextResponse } from "next/server";
import * as jose from "jose";

export async function middleware(request: NextRequest) {
    
    const token = request.cookies.get('token')?.value;
    if (!token) {
        return NextResponse.redirect(new URL("/signin", request.url))
    }
    
    try {
        await jose.jwtVerify(token, Buffer.from("rahasia", 'utf8'))
        return NextResponse.next()
    } catch (error) {
        return NextResponse.redirect(new URL("/signin", request.url))
    }
}
export const config = {
    matcher: ['/explore'],
  }