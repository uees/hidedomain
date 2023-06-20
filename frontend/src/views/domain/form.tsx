import React, { useState } from 'react';
import { Button, Form, Input, Radio, } from 'antd';
import { ActionFunctionArgs, LoaderFunctionArgs, redirect, useLoaderData, useNavigate, useParams } from 'react-router-dom';
import { createDomain, deleteDomain, getDomain, updateDomain } from '../../api/domain';
import { IDomain } from '../../interfaces/models';

export async function deleteDomainAction({ params }: ActionFunctionArgs) {
    await deleteDomain(params.domain as string);
    return redirect("/domains");
}

export async function loader({ params }: LoaderFunctionArgs) {
    const domainData: IDomain = {
        name: '',
        mode: 'local',
        memo: '',
    }

    if (params.domain) {
        //edit route
        const { data } = await getDomain(params.domain)
        Object.assign(domainData, data.data)
    }

    return { data: domainData }
}

const DomainForm: React.FC = () => {

    const { data } = useLoaderData() as { data: IDomain };
    const { domain } = useParams();
    const navigate = useNavigate()

    const onFinish = async (values: IDomain) => {
        if (domain) {
            await updateDomain(domain as string, values)

        } else {
            await createDomain(values)
        }

        return navigate("/domains")
    }

    const [modeValue, setModeValue] = useState<string>(data.mode);
    const onModeChange = ({ mode }: { mode: string }) => {
        if (mode && mode !== modeValue) {
            setModeValue(mode)
        }
    }

    return (
        <Form
            labelCol={{ span: 4 }}
            wrapperCol={{ span: 16 }}
            layout="horizontal"
            style={{ maxWidth: 600 }}
            initialValues={data}
            onValuesChange={onModeChange}
            onFinish={onFinish}
        >
            <Form.Item
                label="域名"
                name='name'
                rules={[{ required: true, message: 'Please input your domain!' }]}
            >
                <Input />
            </Form.Item>
            <Form.Item label="模式" name="mode"
                rules={[{ required: true, message: 'Please input your mode!' }]}
            >
                <Radio.Group
                    options={[{ label: 'Local', value: 'local' }, { label: 'Cloudflare', value: 'cf' },]}
                    optionType="button"
                    buttonStyle="solid"
                />
            </Form.Item>
            {modeValue === 'cf' &&
                <Form.Item label="Token" name='token'>
                    <Input />
                </Form.Item>}
            <Form.Item
                label="备注"
                name='memo'
            >
                <Input.TextArea rows={4} maxLength={6} />
            </Form.Item>

            <Form.Item wrapperCol={{ offset: 4 }}>
                <Button type="primary" htmlType="submit">提交</Button>
            </Form.Item>
        </Form>
    );
};

export default DomainForm;
