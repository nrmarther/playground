import React, {useRef} from 'react';

const Inputref = () =>{

    const inputRef = useRef(null) //reference => {current: null}

    const onClick = () =>{
        console.log('input Value', inputRef.current.value);
    }

    const onClickFocus = () =>{
        inputRef.current.focus();
    }
    return (
        <>
        <input ref={inputRef} />
        <button onClick={onClick}> Log Input Value</button>
        <button onClick={onClickFocus}>Focus On Input</button>
        </>
    )
}
export default Inputref