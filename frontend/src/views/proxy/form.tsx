import React from 'react';
import { Button, Form, Input, } from 'antd';
import { ActionFunctionArgs, LoaderFunctionArgs, redirect, useLoaderData, useNavigate, useParams } from 'react-router-dom';
import { createProxyItem, deleteProxyItem, getProxyItem, updateProxyItem } from '../../api/proxyitem';
import { IProxyItem } from '../../interfaces/models';
import { useTitle } from '../../hooks';

export async function deleteProxyItemAction({ params }: ActionFunctionArgs) {
    await deleteProxyItem(params.id as string);
    return redirect("/proxies");
}

export async function loader({ params }: LoaderFunctionArgs) {
    const proxyItem: IProxyItem = {
        content: '',
        protocol: '',
        memo: '',
    }

    if (params.id) {
        //edit route
        const { data } = await getProxyItem(params.id)
        Object.assign(proxyItem, data.data)
    }

    return { data: proxyItem }
}

const ProxyItemForm: React.FC = () => {
    useTitle("添加/编辑代理项");
    const { data } = useLoaderData() as { data: IProxyItem };
    const { id } = useParams();
    const navigate = useNavigate()

    const onFinish = async (values: IProxyItem) => {
        if (id) {
            await updateProxyItem(id as string, values)

        } else {
            await createProxyItem(values)
        }

        return navigate("/proxies")
    }

    return (
        <Form
            labelCol={{ span: 4 }}
            wrapperCol={{ span: 16 }}
            layout="horizontal"
            style={{ maxWidth: 600 }}
            initialValues={data}
            onFinish={onFinish}
        >
            <Form.Item
                label="协议"
                name='protocol'
                rules={[{ required: true, message: 'Please input your protocol!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="标题"
                name='memo'
                rules={[{ required: true, message: 'Please input your memo!' }]}
            >
                <Input />
            </Form.Item>

            <Form.Item
                label="内容"
                name='content'
                rules={[{ required: true, message: 'Please input your content!' }]}
            >
                <Input.TextArea rows={20} />
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 4 }}>
                <Button type="primary" htmlType="submit">提交</Button>
            </Form.Item>
        </Form>
    );
};

export default ProxyItemForm;
