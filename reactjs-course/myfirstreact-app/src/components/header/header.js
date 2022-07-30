import React from 'react';

const Header = () =>{
    //allows to dynamically change className's using variables instead of hardcoded values. can use var, let, const
    const logoTitle = 'My First React App';
    return (
        <header className='header row mt-3'>
            <div className='logo col-4'>
                {/* calls the previously declared variable logoTitle in curly braces */}
                <a href="#">{logoTitle}</a>
            </div>
            {/* more classNames from bootstrap */}
            <div className='col-4'></div>
            {/* multiple classNames used on a single element */}
            <div className='col-4 search'>
                <div className='input-group'>
                    <input type="text" className="form-control" placeholder="Search Keyword" />
                    <div className='input-group-append'>
                        <button className='btn btn-outline-secondary'>Search </button>
                    </div>
                </div>
            </div>
        </header>
    )
}
export default Header;