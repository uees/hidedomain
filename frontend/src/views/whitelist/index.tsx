/* eslint-disable jsx-a11y/anchor-is-valid */
import React from 'react';
import { ActionFunctionArgs, LoaderFunctionArgs, redirect, useLoaderData, useNavigate, useParams, useSubmit } from "react-router-dom";
import { Button, Popconfirm, Table, message } from 'antd';
import { addIPRule, deleteIPRule, showList, updateIPRule } from "../../api/whitelist";
import { EditableRow, EditableCell } from './row'
import { getIP } from '../../api/ip';
import './style.css';
import { useTitle } from '../../hooks';

export interface IPRuleDataType {
    key?: string | number;
    domain?: string;
    ip?: string;
    memo?: string;
}

export async function loader({ params }: LoaderFunctionArgs) {
    const { data } = await showList(params.domain as string)
    const ipRules = data.data as IPRuleDataType[];
    return { ipRules }
}

export async function deleteRuleAction({ params }: ActionFunctionArgs) {
    const { ruleid, domain } = params as { ruleid: string, domain: string }
    await deleteIPRule(domain, ruleid)
    return redirect(`/domains/${domain}/whitelist`)
}

type EditableTableProps = Parameters<typeof Table>[0];

type ColumnTypes = Exclude<EditableTableProps['columns'], undefined>;

const App: React.FC = () => {
    useTitle("白名单列表");
    const { ipRules } = useLoaderData() as { ipRules: IPRuleDataType[] }
    const { domain } = useParams() as { domain: string };
    const submit = useSubmit();
    const navigate = useNavigate();

    const defaultColumns: (ColumnTypes[number] & { editable?: boolean; dataIndex: string })[] = [
        {
            title: '域名',
            dataIndex: 'domain',
        },
        {
            title: 'IP',
            dataIndex: 'ip',
        },
        {
            title: '备注',
            dataIndex: 'memo',
            editable: true,
            width: '30%',
        },
        {
            title: '操作',
            dataIndex: 'operation',
            render: (_, record: IPRuleDataType) => (
                <Popconfirm title="Sure to delete?"
                    onConfirm={() => submit(null, { method: 'delete', action: `/domains/${record.domain}/whitelist/${record.key}/destroy` })}
                >
                    <a>Delete</a>
                </Popconfirm>
            ),
        },
    ];

    const handleAdd = async () => {
        const { data } = await getIP();
        const { ip, network, version } = data
        // console.log(ip, network, version)
        if (version === "IPv4") {
            await addIPRule(domain, { ip: network, memo: ip })
            // message.success("add success");
            return navigate(`/domains/${domain}/whitelist`, { replace: true })
        } else {
            message.error("please disable IPV6");
        }
    };

    const handleSave = async (row: IPRuleDataType) => {
        await updateIPRule(domain, row.key as string, { memo: row.memo })
        message.success("update success");
        return navigate(`/domains/${domain}/whitelist`, { replace: true })
    };

    const components = {
        body: {
            row: EditableRow,
            cell: EditableCell,
        },
    };

    const columns = defaultColumns.map((col) => {
        if (!col.editable) {
            return col;
        }
        return {
            ...col,
            onCell: (record: IPRuleDataType) => ({
                record,
                editable: col.editable,
                dataIndex: col.dataIndex,
                title: col.title,
                handleSave,
            }),
        };
    });

    return (
        <div>
            <Button onClick={handleAdd} type="primary" style={{ marginBottom: 16 }} size='large'>
                Add a rule
            </Button>
            <Table
                components={components}
                rowClassName={() => 'editable-row'}
                dataSource={ipRules}
                columns={columns as ColumnTypes}
                pagination={false}
            />
        </div>
    );
};

export default App;
