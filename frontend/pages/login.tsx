import React, { Component } from 'react'

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
                    <a href="/register" style={{color: "#1955c8", textDecoration: "underline"}}>Don&apos;t have an account?</a>
                </div>
            </div>
        )
    }
}
