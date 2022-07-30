import React, {useState} from "react";
import axios from "axios";

const Adduser = () =>{
    const [name, setName] = useState();
    const handlechange = (e) =>{
        setName(e.target.value);
    }

    const handleSubmit = (event) =>{
        event.preventDefault();
        const apiURL = "https://jsonplaceholder.typicode.com/users";
        const user = {
            name: name,
            phone: "000-000-0000",
            username: "Lobcobob",
            website: "website.web.com"
        }

        axios.post(apiURL, user)
        .then(res => {
            console.log(res);
        })
    }
    return(
        <>
            <form onSubmit={(event) => handleSubmit(event)}>
                <label>Person Name</label>
                <input type="text" name="name" value={name} onChange={(event) => handlechange(event)} />
                <button type="submit">Add</button>
            </form>
        </>
    )
}
export default Adduser;