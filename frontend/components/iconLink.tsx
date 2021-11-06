import React from "react"
import styles from "../styles/index.module.css"

interface Props {
    text: string;
    href?: string;
    style?: React.CSSProperties;
}

export default function iconLink (props: React.PropsWithChildren<Props>) {
    return (
        <div className={styles["icon-link"]} style={props.style}>
            {props.children}{' '}<a className={styles["nav-boldlink"]} href={props.href}>{props.text}</a>
        </div>
    )
}
