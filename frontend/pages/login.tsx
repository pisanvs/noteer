/**
* Copyright (C) 2021  Maximiliano Morel (pisanvs) <maxmorel@pisanvs.cl>
*
* This file is part of Noteer, a note taking application.
* 
* Noteer is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License v3 as
* published by the Free Software Foundation
*
* Noteer is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with Noteer.  If not, see <https://www.gnu.org/licenses/>.
*
*
* @license GPL-3.0 <https://www.gnu.org/licenses/gpl-3.0.txt>
*/

import React, { Component } from 'react'
import Link from 'next/link'

interface Props {
	
}
interface State {
	
}

export default class login extends Component<Props, State> {
	state = {}

	render() {
		return (
			<div className="login-main">
				<div className="login-form-wrapper">

					<h2 style={{justifySelf: 'center'}}>Login</h2>

					<form action="https://api.noteer.local/api/v1/auth/login" method="POST">
						<label htmlFor="username">Username:</label>
						<br />
						<input type="text" name="username" id="username" className="login-input"/>

						<br />
						<br />

						<label htmlFor="password">Password:</label>
						<br />
						<input type="password" name="password" id="password" className="login-input" />

						<br />
						<input type="submit" name="submit" id="submit" value="Submit" className="login-input"/>
					</form>
					<Link href="/register" passHref><a style={{color: "#1955c8", textDecoration: "underline"}}>Don&apos;t have an account?</a></Link>
				</div>
			</div>
		)
	}
}
