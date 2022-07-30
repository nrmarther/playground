import React from 'react';

const Main = () =>{
    return (
        <div className="main-content">
            Main Content goes here
            {/* classNames from bootstrap */}
            <div className="row">
                <div className='col-6'>
                    Column1
                </div>
                <div className='col-6'>
                    Column2
                </div>
            </div>
        </div>
    )
}
export default Main;