import React, {userState, useEffect, useState} from "react";
import { userDetail } from "./userdetail"

const Todo = () => {
    const [userDetail, setUserDetail] = useState(false);
    const [flag, setFlag] = useState(false);
    // I have no idea why i need this function but when I call setFlag() on ln23 directly, it breaks the code.
    const fetchUserData = () =>{
        setFlag(true);
    }
    useEffect(() =>  {
        // makes sure the api call is only made once
        if(flag){
            const apiURL = "https://jsonplaceholder.typicode.com/users"
            fetch(apiURL).then((response) => response.json())
            .then((json) => setUserDetail(json))
            .then((e) => console.log(e));
            //resets flag to false after the first call. without this line it will make the api call infinitely
            setFlag(false);
        }
    },[flag]);
    return (
        <>
            <button className="btn btn-primary" onClick={fetchUserData}>Fetch User Data</button>
            {userDetail && userDetail.map((user) => (
                <div className="row">
                    <div className="col-3">{user.username}</div>
                    <div className="col-3">{user.name}</div>
                    <div className="col-3">{user.phone}</div>
                    <div className="col-3">{user.website}</div>
                </div>
            ))}
        </>
    );
}
export default Todo;