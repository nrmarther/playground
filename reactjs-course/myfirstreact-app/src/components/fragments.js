/*
react does not allow multiple elements (such as divs) to be passed
used in the return statement. They must all be wrapped into a div OR 
you can use a React Fragment. React Fragments do not show up when 
inspecting the element in browser satisfy the requirement to not have
multiple elements in the return statement of the functional component. 
React Fragements can be used as seen below or with empty carots as shown below

<>
    <div>
    </div>
    <div>
    </div>
</>
*/
import React from "react";
const Fragments = () => {
    return (
        <React.Fragment>
            <div>
                To Do Column 1
            </div>
            <div>
                To Do Column 2
            </div>
        </React.Fragment>
    )
}
export default Fragments;