/* eslint-disable jsx-a11y/anchor-is-valid */
import React from 'react';
import { LoaderFunctionArgs, useLoaderData, useNavigate, useSubmit } from "react-router-dom";
import { Button, Space, Table } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import type { ColumnsType } from 'antd/es/table';
import { IProxyItem } from '../../interfaces/models';
import { useTitle } from '../../hooks';
import { proxiesList } from '../../api/proxyitem';

export async function loader({ request }: LoaderFunctionArgs) {
    const { data } = await proxiesList()

    // add key, fix antd warning
    let proxies = data.data as IProxyItem[];
    for (let proxyItem of proxies) {
        proxyItem.key = proxyItem.id
    }
    return { proxies }
}

const Proxyitem: React.FC = () => {
    useTitle("代理项管理");
    const { proxies } = useLoaderData() as { proxies: IProxyItem[] };
    const navigate = useNavigate();
    const submit = useSubmit()

    function handleAddProxyItem() {
        return navigate("/proxies/add")
    }

    const columns: ColumnsType<IProxyItem> = [
        {
            title: '协议',
            dataIndex: 'protocol',
            key: 'protocol',
            render: (_, { id, protocol }) => (
                <a onClick={() => {
                    submit(null, { method: "get", action: `/proxies/${id}/edit` })
                }}>{protocol}</a>
            ),
        },
        {
            title: '标题',
            dataIndex: 'memo',
            key: 'memo',
            render: (_, { id, memo }) => {
                return (
                    <a onClick={() => {
                        submit(null, { method: "get", action: `/proxies/${id}/edit` })
                    }}>{memo}</a>
                )
            },
        },
        {
            title: '内容',
            key: 'content',
            dataIndex: 'content',
        },
        {
            title: 'Action',
            key: 'action',
            render: (_, record) => (
                <Space size="middle">
                    <a onClick={() => {
                        // eslint-disable-next-line no-restricted-globals
                        if (confirm("Please confirm you want to delete this record.")) {
                            submit(null, { method: "delete", action: `/proxies/${record.id}/destroy` })
                        }
                    }}>Delete</a>
                </Space >
            ),
        },
    ];

    return (
        <>
            <Button type="primary"
                icon={<PlusOutlined />}
                size='large'
                style={{ marginBottom: 16 }}
                onClick={handleAddProxyItem}>添加</Button>
            <Table columns={columns} dataSource={proxies} />
        </>

    )
}

export default Proxyitem;
