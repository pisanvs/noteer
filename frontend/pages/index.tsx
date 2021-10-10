import React, { Component } from "react"
import Head from "next/head"
import IconLink from "../components/iconLink"
import { AiOutlineUser } from "@react-icons/all-files/ai/AiOutlineUser"
import { IoHelp } from "@react-icons/all-files/io5/IoHelp"
import Image from "next/image"

interface Props {

}
interface State {

}

const iconLinkStyle = {

}

const iconStyle: React.CSSProperties = {
	display: "inline-flex",
	margin: "1px 4px",
	width: "1em",
	height: "1em"
}

const logoStyle: React.CSSProperties = {
	margin: "10px 5px",
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
						<Image src="/placeholder-logo-1.svg" width="120" height="30" alt="logo" className="header-logo"/>
						<IconLink text="Login"><AiOutlineUser style={iconStyle} /></IconLink>
						<IconLink text="Help"><IoHelp style={iconStyle} /></IconLink>
					</div>
					<div className="side side-left">
						
					</div>
				</header>
				<div className="main container">

				</div>
				<div className="footer container"></div>
			</>
		)
	}
}
