import React, {useState} from 'react';

import { Tabs, TabPanel } from 'react-tabs';
import 'react-tabs/style/react-tabs.css';
import PageTitle from '../components/pagetitle';
import { useAuth } from '../AuthContext';
import { useNavigate } from "react-router-dom";

import {Link} from 'react-router-dom';

Login.propTypes = {
    
};

function Login(props) {
    const [formData, setFormData] = useState({
        email: '',
        password: ''
    });

    const handleChange = (e) => {
        setFormData({
            ...formData,
            [e.target.id]: e.target.value
        });
    };

    const [errors, setErrors] = useState({});
    const navigate = useNavigate();
    const { login } = useAuth();

    const handleSubmit = async (e) => {
        e.preventDefault();

        console.log(formData)
        const errors = {};
        if (!formData.email.includes('@')) {
            errors.email = 'Please enter a valid email address.';
        }
        if (formData.password.trim() === '') {
            errors.password = 'Password is required.';
        }

        if (Object.keys(errors).length === 0) {
            localStorage.setItem("token", "123");
            login()
        
            try {
                const response = await fetch("/api/user/login", {
                    method: "POST",
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(formData)
                });
    
                if (!response.ok) {
                    throw new Error("Network response was not ok");
                }
    
                navigate("/");
            } catch (error) {
                console.error(error);
            }
        }
            
        setErrors(errors);
    };


    return (
        <div>
            <PageTitle heading='Login' title='Login' />

            <section className="register login">
            <div className="container">
                <div className="row">
                <div className="col-md-12">
                    <div className="block-text center">
                    <h3 className="heading">Login To Rockie</h3>
                    <p className="desc fs-20">
                        Welcome back! Log In now to start trading
                    </p>
                    </div>
                </div>
                <div className="col-md-12">
                <Tabs>
                    <TabPanel>
                        <div className="content-inner">
                            <form onSubmit={handleSubmit}>
                                <div className="form-group">
                                <label for="email">Email</label>
                                <input
                                    type="email"
                                    className="form-control"
                                    id="email"
                                    placeholder="Please fill in the email form."
                                    autoComplete="username"
                                    onChange={handleChange}
                                    required
                                />
                                {errors.email && <p className="error">{errors.email}</p>}
                                </div>
                                <div className="form-group s1">
                                <label>Password </label>
                                <input
                                    type="password"
                                    className="form-control"
                                    id="password"
                                    placeholder="Please enter a password."
                                    autoComplete="new-password"
                                    onChange={handleChange}
                                    required
                                />
                                {errors.password && <p className="error">{errors.password}</p>}
                                </div>

                                <button type="submit" className="btn-action">Login</button>
                                <div className="bottom">
                                <p>Not a member?</p>
                                <Link to="/register">Register</Link>
                                </div>
                            </form>
                        </div>
                    </TabPanel>                    

                </Tabs> 
                </div>
                </div>
            </div>
            </section>
            
        </div>
    );
}

export default Login;