import React, { Component } from 'react'

interface Props {

}
interface State {
    isValid: boolean;
    passwordState: string;
}

const allowedPasswordsChars: string[] = [
    'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z', 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '!', '@', '#', '$', '%', '^', '&', '*', '(', ')', '-', '_', '+', '=', '{', '[', '}', ']', '|', ':', ';', '"', '\'', '<', ',', '>', '.', '?', '/', '~', '`'
];

export default class register extends Component<Props, State> {
    passwordConfirmation: React.RefObject<HTMLInputElement>;
    password: React.RefObject<HTMLInputElement>;
    constructor(props: Props) {
        super(props)
        this.state = {
            isValid: false,
            passwordState: ""
        }
        this.password = React.createRef();
        this.passwordConfirmation = React.createRef();
    }

    onPasswordUpdate(e: React.ChangeEvent<HTMLInputElement>) {
        // check if password is too short or too long
        if (this.password.current?.value?.length !== undefined && this.password.current.value.length < 8) {
            this.setState({ isValid: false, passwordState: "Password is too short" })
        } else if (this.password.current?.value?.length !== undefined && this.password.current.value.length > 64) {
            this.setState({ isValid: false, passwordState: "Password is too long" })
        }

        // check if password contains invalid characters
        this.password.current?.value?.split("").forEach(char => {
            if (!allowedPasswordsChars.includes(char)) {
                this.setState({ isValid: false, passwordState: "Password contains invalid characters" })
            }
        })

        // Check if both passwords match
        this.password.current?.value === this.passwordConfirmation.current?.value ? this.setState({ isValid: true, passwordState: "" }) : this.setState({ isValid: false, passwordState: "Passwords do not match" })
    }

    render() {
        return (
            <div className="register-main">
                <div className="register-form-wrapper">

                    <h2 style={{ justifySelf: 'center' }}>Register</h2>

                    <form action="http://api.noteer.local/api/v1/auth/register" method="post">
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
                        <input type="password" name="password" id="password" ref={this.password} className="register-input" />

                        <br />
                        <br />

                        <label htmlFor="passwordConfirmation">Confirm Password:</label>
                        <br />
                        <input type="password" name="passwordConfirmation" id="passwordConfirmation" ref={this.passwordConfirmation} className="register-input" onChange={(e) => this.onPasswordUpdate(e)} />

                        <br />

                        {
                            this.state.passwordState !== "" ? <p className="register-error">{this.state.passwordState}</p> : null
                        }

                        <br />
                        <br />


                        <input type="submit" name="submit" id="submit" value="Submit" className="register-input" disabled={!this.state.isValid} />

                        <br />
                        <br />

                    </form>
                </div>
            </div>
        )
    }
}
