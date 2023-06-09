/* eslint-disable jsx-a11y/anchor-is-valid */
import React from 'react';
import { Link, LoaderFunctionArgs, useLoaderData, useNavigate, useSubmit } from "react-router-dom";
import { Button, Space, Table, Tag } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import type { ColumnsType } from 'antd/es/table';
import { IDomain } from '../../interfaces/models';
import { useTitle } from '../../hooks';
import { domainList } from '../../api/domain';

export async function loader({ request }: LoaderFunctionArgs) {
    const { data } = await domainList()

    // 添加 key fix antd error
    let domains = data.data as IDomain[];
    for (let domain of domains) {
        domain.key = domain.id
    }
    return { domains }
}

const Domain: React.FC = () => {

    useTitle("域名管理");
    const { domains } = useLoaderData() as { domains: IDomain[] };
    const navigate = useNavigate();
    const submit = useSubmit()

    function handleAddDomain() {
        return navigate("/domains/add")
    }

    const columns: ColumnsType<IDomain> = [
        {
            title: '域名',
            dataIndex: 'name',
            key: 'name',
            render: (_, { name }) => (
                <a onClick={() => {
                    submit(null, { method: "get", action: `/domains/${name}/edit` })
                }}>{name}</a>
            ),
        },
        {
            title: '模式',
            dataIndex: 'mode',
            key: 'mode',
            render: (_, { mode }) => {
                let color = 'geekblue'
                if (mode === 'local') {
                    color = "magenta"
                }
                return (
                    <Tag color={color} key={mode}>
                        {mode.toUpperCase()}
                    </Tag>
                )
            },
        },
        {
            title: '备注',
            key: 'memo',
            dataIndex: 'memo',
        },
        {
            title: 'Action',
            key: 'action',
            render: (_, record) => (
                <Space size="middle">
                    {record.mode === 'local' ?
                        <Link to={`/domains/${record.name}/whitelist`}>IP List</Link> :
                        <>
                            <a>CloudflareRules</a>
                            <a>Token</a>
                        </>}
                    <a onClick={() => {
                        // eslint-disable-next-line no-restricted-globals
                        if (confirm("Please confirm you want to delete this record.")) {
                            submit(null, { method: "delete", action: `/domains/${record.name}/destroy` })
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
                onClick={handleAddDomain}>添加</Button>
            <Table columns={columns} dataSource={domains} />
        </>

    )
}

export default Domain;
