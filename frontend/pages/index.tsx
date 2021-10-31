import React, { Component } from "react"
import Head from "next/head"
import IconLink from "../components/iconLink"
import { AiOutlineUser } from "@react-icons/all-files/ai/AiOutlineUser"
import { IoHelp } from "@react-icons/all-files/io5/IoHelp"
import { AiOutlineArrowUp } from "@react-icons/all-files/ai/AiOutlineArrowUp"
import Image from "next/image"
import Link from "next/link"

interface Props {

}
interface State {

}

const iconStyle: React.CSSProperties = {
	display: "inline-flex",
	margin: "1px 4px",
	minWidth: "1em",
	minHeight: "1em"
}

export default class index extends Component<Props, State> {
	state = {}

	render() {
		return (
			<>
				<Head>
					<title>Noteer, notes but better</title>
				</Head>
				<header className="header">
					<div className="side side-left">
						<Image 
							src="/placeholder-logo-1.svg"
							width="120"
							height="30"
							alt="logo"
							className="header-logo"
						/>
						
						<Link href="/login" passHref>
							<IconLink text="Login"><AiOutlineUser style={iconStyle} /></IconLink>
						</Link>

						<Link href="/help" passHref>
							<IconLink text="Help"><IoHelp style={iconStyle} /></IconLink>
						</Link>

						<Link href="/app" passHref>
							<IconLink text="Access"><AiOutlineArrowUp style={iconStyle} /></IconLink>
						</Link>

					</div>
					<div className="side side-right">
						<Link href="/register">
							<a style={{margin: "12px 10px", fontWeight:300} as React.CSSProperties} className="link">Register</a>
						</Link>
						<Link href="#about">
							<a style={{margin: "12px 10px", fontWeight:300} as React.CSSProperties} className="link">About Us</a>
						</Link>
						<Link href="https://github.com/pisanvs/noteer">
							<a style={{margin: "12px 10px", fontWeight:300} as React.CSSProperties} className="link">Github Repo</a>
						</Link>
					</div>
				</header>
				<div className="main container">

				</div>
				<div className="footer container"></div>
			</>
		)
	}
}
