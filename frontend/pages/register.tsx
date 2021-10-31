import React, { Component } from 'react'

interface Props {

}
interface State {
    isValid: boolean;
}

export default class register extends Component<Props, State> {
    constructor(props: Props) {
        super(props)
        this.state = {
            isValid: false
        }
    }

    onPasswordUpdate() {

    }

    render() {
        return (
            <div className="register-main">
                <div className="register-form-wrapper">

                    <h2 style={{ justifySelf: 'center' }}>Register</h2>

                    <form action="noteer.local/api/v1/auth/register" method="post">
                        <label htmlFor="name">Name:</label>
                        <br />
                        <input type="text" name="name" id="name" className="register-input" />

                        <br />
                        <br />

                        <label htmlFor="username">Username:</label>
                        <br />
                        <input type="text" name="username" id="username" className="register-input" />

                        <br />
                        <br />

                        <label htmlFor="email">Email:</label>
                        <br />
                        <input type="email" name="email" id="email" className="register-input" />

                        <br />
                        <br />

                        <label htmlFor="password">Password:</label>
                        <br />
                        <input type="password" name="password" id="password" className="register-input" />

                        <br />
                        <br />

                        <label htmlFor="password">Password:</label>
                        <br />
                        <input type="password" name="password" id="password" className="register-input" />

                        <br />
                        <input type="submit" name="submit" id="submit" value="Submit" className="register-input" disabled={!this.state.isValid} />
                    </form>
                </div>
            </div>
        )
    }
}
