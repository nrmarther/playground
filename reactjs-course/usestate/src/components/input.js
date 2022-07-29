import React from "react";

const Input = ({input, setInput, placeholder}) => {
    return(
        <>
            <input type="text" placeholder={placeholder} value={input} onChange={(event) => setInput(event.target.value)} />
        </>
    );
}
export default Input;