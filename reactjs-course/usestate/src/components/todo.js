import React, {useState} from "react";
import Input from "./input";

const Todo = () => {
    const [name, setName] = useState();
    const [city, setCity] = useState();
    return (
        <>
            <Input 
                input={name}
                setInput={(value)=> setName(value)}
                placeholder="Enter your name"
            />
            <br />
            {/* shows no value when text-field is not occupied */}
            {name &&
               <>Name: {name}</>
            }
            <br />
            <Input 
                input={city}
                setinput={(value)=> setCity(value)}
                placeholder="Enter your city"
            />
            <br/>
            {/* gives default message when text-field is not occupied */}
            {city ?
            <> City: {city} </>
            :
            'No city found'
            }
            <br/>


        </>
    );
}
export default Todo;