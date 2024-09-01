import React, { useState } from "react";
import { Navigate, useLocation, useNavigate } from "react-router-dom";
import { Form, Input, Button, message, Spin } from "antd";
import { useStore, useTitle } from "../../hooks";
import "../../styles/login.css";

const Login: React.FC = () => {

    useTitle('Login')
    const { user } = useStore()
    const [loading, setLoading] = useState(false);
    const navigate = useNavigate();

    //useEffect(() => {
    //    site.setMenu([{ label: '登录', key: '/login', }])
    //})

    const handleLogin = async (username: string, password: string) => {
        setLoading(true);
        try {
            await user.login({ username, password })
            await user.loadInfo()
            message.success("Login success");
            navigate("/");
        } catch (e: any) {
            setLoading(false);
            message.error(e);
        }
    };

    const onFinish = async (values: any) => {
        // console.log('Success:', values);
        const { username, password } = values;
        await handleLogin(username, password);
    };

    const onFinishFailed = (errorInfo: any) => {
        //console.log('Failed:', errorInfo);
    };

    const location = useLocation();

    if (user.token) {
        // redirect
        return <Navigate replace to='/' state={{ location }} />;
    }

    return (
        <div className="login-container">
            <Form name="basic"
                className="content"
                labelCol={{ span: 8 }}
                wrapperCol={{ span: 16 }}
                style={{ maxWidth: 600, paddingTop: 40 }}
                initialValues={{ remember: true }}
                onFinish={onFinish}
                onFinishFailed={onFinishFailed}
                autoComplete="off">

                <Spin spinning={loading} tip="Login...">
                    <Form.Item
                        label="Username"
                        name="username"
                        rules={[{ required: true, message: 'Please input your username!' }]}
                    >
                        <Input />
                    </Form.Item>

                    <Form.Item
                        label="Password"
                        name="password"
                        rules={[{ required: true, message: 'Please input your password!' }]}
                    >
                        <Input.Password />
                    </Form.Item>

                    <Form.Item wrapperCol={{ offset: 8, span: 16 }}>
                        <Button type="primary" htmlType="submit">
                            Submit
                        </Button>
                    </Form.Item>
                </Spin>
            </Form>
        </div>
    );
};

export default Login;
