import React, { ReactElement } from "react"

interface Props {
    text: string;
}

const style: React.CSSProperties = {
    display: "inline-flex",
    margin: "10px 10px 10px 20px",
    justifySelf: "center",
    width: "4em",
}

export default function iconLink (props: React.PropsWithChildren<Props>) {
    return (
        <div className="icon-link" style={style}>
            {props.children}{' '}<span className="il-text">{props.text}</span>
        </div>
    )
}
