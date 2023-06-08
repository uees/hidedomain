/* eslint-disable jsx-a11y/anchor-is-valid */
import React, { useEffect } from 'react';
import { Link, LoaderFunctionArgs, useLoaderData } from "react-router-dom";
import { Button, Space, Table, Tag } from 'antd';
import { PlusOutlined } from '@ant-design/icons';
import type { ColumnsType } from 'antd/es/table';
import { IDomain } from '../../interfaces/models';
import { useStore, useTitle } from '../../hooks';
import { domainList } from '../../api/domain';

const columns: ColumnsType<IDomain> = [
    {
        title: '域名',
        dataIndex: 'name',
        key: 'name',
        render: (text) => <a>{text}</a>,
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
                <Link to={`/domain/${record.name}/whitelist`}>IP List</Link>
                <a>Delete</a>
            </Space>
        ),
    },
];

export async function loader({ request }: LoaderFunctionArgs) {
    const { data } = await domainList()
    return data
}

const Domain: React.FC = () => {

    useTitle("域名管理");

    const { site } = useStore()

    useEffect(() => {
        site.setBreadcrumb([{ title: '主页' }, { title: '域名' }]);
    })

    const { data } = useLoaderData() as { data: IDomain[] }

    return (
        <>
            <Button type="primary"
                icon={<PlusOutlined />}
                size='large'
                style={{ marginTop: '8px', marginLeft: '8px' }}>添加</Button>
            <Table columns={columns} dataSource={data} />
        </>

    )
}

export default Domain;
