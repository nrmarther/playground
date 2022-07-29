import React, {useState} from "react";
import Input from "./input";
import { userDetail } from "./userdetail"

const Todo = () => {
    const [name, setName] = useState();
    const [city, setCity] = useState();
    console.log('userDetail', userDetail)
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
            <br/><br />
            {userDetail.map((user) => (
                <div className="row">
                    <div className="col-4">{user.name}</div>
                    <div className="col-4">{user.salary}</div>
                    <div className="col-4">{user.job}</div>
                    <br />
                </div>
            ))}
            

        </>
    );
}
export default Todo;