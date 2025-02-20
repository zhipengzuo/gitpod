/**
 * Copyright (c) 2023 Gitpod GmbH. All rights reserved.
 * Licensed under the GNU Affero General Public License (AGPL).
 * See License.AGPL.txt in the project root for license information.
 */

import { OrganizationSettings } from "@gitpod/gitpod-protocol";
import { Entity, Column, PrimaryColumn } from "typeorm";
import { TypeORM } from "../typeorm";

@Entity()
export class DBOrgSettings implements OrganizationSettings {
    @PrimaryColumn(TypeORM.UUID_COLUMN_TYPE)
    orgId: string;

    @Column({
        default: false,
    })
    workspaceSharingDisabled?: boolean;

    @Column("varchar", { nullable: true })
    defaultWorkspaceImage?: string | null;

    @Column("json", { nullable: true })
    allowedWorkspaceClasses?: string[] | null;

    @Column()
    deleted: boolean;
}
