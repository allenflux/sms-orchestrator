import json
import os
import time

import click

from sms_client.sms_core import m_list_project, m_create_project, m_list_device_main, m_allocate_device_2_project, \
    m_list_group, m_create_group, m_list_device_sub, m_allocate_device_2_group, m_sub_upload_task, m_list_tasks, \
    m_list_tasks_sub, m_download_task_file, decode_download_task_file, m_controller_main_sms_record_list, \
    m_controller_sub_sms_record_list, m_sub_get_conversation_record_list, m_main_get_conversation_record_list


@click.group()
def sms_cli():
    """sms_cli"""
    pass

@click.command(name="list-device")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
def list_device(page):
    """list-device-main """
    out = m_list_device_main(page)
    click.echo(out)

@click.command(name="sub-list-device")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub-user-id",
    type=int,
    required=True,
)
def sub_list_device(page, sub_user_id):
    """list-device-sub """
    out = m_list_device_sub(page,sub_user_id)
    click.echo(out)

# Project
@click.command(name="list-project")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
def list_project(page):
    """list-project"""
    out = m_list_project(page)
    click.echo(out)

@click.command(name="create-project")
@click.option(
    "-project-name",
    "--project-name",
    help="project name",
    type=str,
    required=True,
)
@click.option(
    "-note",
    "--note",
    help="note",
    default="Create By (sms cli)",
    type=str,
)
def create_project(project_name, note):
    """create-project"""
    out = m_create_project(project_name=project_name, note=note)
    click.echo(out)

@click.command(name="delete-project")
def delete_project():
    """delete-project"""
    pass
@click.command(name="update-projec")
def update_project():
    """update-project"""
    pass

@click.command(name="allocate-device-to-project")
@click.option(
    "-device-id",
    "--device-id",
    help="device-id",
    type=int,
    required=True,
)
@click.option(
    "-project-id",
    "--project-id",
    help="project-id",
    type=int,
    required=True,
)
def allocate_device_to_project(device_id,project_id):
    """allocate-device-to-project"""
    out = m_allocate_device_2_project(device_id, project_id)
    click.echo(out)

# Task
@click.command(name="create-task")
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub-user-id",
    type=int,
    required=True,
)
@click.option(
    "-f",
    "--file",
    help="file path",
    type=click.Path(exists=True),
    required=True,
)
@click.option(
    "-task-name",
    "--task-name",
    help="task name",
    type=str,
    required=True,
)
@click.option(
    "-group-id",
    "--group-id",
    help="group-id",
    type=int,
    required=True,
)
@click.option(
    "-interval-time",
    "--interval-time",
    help="interval-time ",
    type=int,
    default=1,
    required=False,
)
@click.option(
    "-timing-start-time",
    "--timing-start-time",
    help="timing-start-time ",
    type=str,
    default=str(time.time()),
    required=False,
)
def create_task(sub_user_id,file, task_name, group_id, interval_time, timing_start_time):
    """create-task"""
    out = m_sub_upload_task(
        dict(
            task_name=task_name,
            group_id=group_id,
            sub_user_id=sub_user_id,
            file=file,
            timing_start_time=timing_start_time,
            interval_time=str(interval_time),
        )
    )
    click.echo(out)
@click.command(name="delete-task")
def delete_task():
    """delete-task"""
    pass

@click.command(name="list-tasks")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
def list_tasks(page):
    """list-tasks"""
    out = m_list_tasks(page)
    click.echo(out)

@click.command(name="sub-list-tasks")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="Sub User ID",
    type=int,
    required=True,
)
def sub_list_tasks(page, sub_user_id):
    """list-tasks"""
    out = m_list_tasks_sub(page, sub_user_id)
    click.echo(out)

@click.command(name="download-task")
@click.option(
    "-file-name",
    "--file-name",
    help="file name",
    type=str,
    required=True,
)
def download_task(file_name):
    """download-task"""
    data = m_download_task_file(file_name)
    if data["code"] != 0:
        click.echo(data)
        return
    click.echo("Interpret the contents of the file.")
    out, dict_data = decode_download_task_file(data)
    click.echo(out)
    file_path = os.getcwd() + "/" + file_name
    click.echo("File storage path:" + file_path)
    with open(file_path, "w") as f:
        json.dump(dict_data, f)


@click.command(name="list-task-record")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
def list_task_record(page):
    """list-task-record"""
    out = m_controller_main_sms_record_list(page)
    click.echo(out)

@click.command(name="sub-list-task-record")
@click.option(
    "-p",
    "--page",
    default=1,
    help="List Page Num",
)
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub user id",
    type=int,
    required=True,
)
def sub_list_task_record(page, sub_user_id):
    """sub-list-task-record"""
    out = m_controller_sub_sms_record_list(page, sub_user_id)
    click.echo(out)

# Group
@click.command(name="list-groups")
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub user id",
    type=int,
    required=True,
)
def list_groups(sub_user_id):
    """list-groups"""
    out = m_list_group(sub_user_id)
    click.echo(out)

@click.command(name="create-group")
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub-user-id",
    type=int,
    required=True,
)
@click.option(
    "-project-id",
    "--project-id",
    help="project-id",
    type=int,
    required=True,
)
@click.option(
    "-group-name",
    "--group-name",
    help="group name",
    type=str,
    required=True,
)
def create_group(sub_user_id, project_id, group_name):
    """create-group"""
    out = m_create_group(group_name=group_name, project_id=project_id, sub_user_id=sub_user_id)
    click.echo(out)

@click.command(name="allocate-device-to-group")
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub-user-id",
    type=int,
    required=True,
)
@click.option(
    "-group-id",
    "--group-id",
    help="group-id",
    type=int,
    required=True,
)
@click.option(
    "-device-id",
    "--device-id",
    help="device-id",
    type=int,
    required=True,
)
def allocate_device_to_group(sub_user_id, group_id, device_id ):
    """allocate-device-to-group"""
    out = m_allocate_device_2_group(sub_user_id=sub_user_id, group_id=group_id, device_id=device_id)
    click.echo(out)
    pass
@click.command(name="update-group")
def update_group():
    """update-group"""
    pass
@click.command(name="delete-group")
def delete_group():
    """delete-group"""
    pass
# Chat
@click.command(name="list-chats")
@click.option(
    "-project-id",
    "--project-id",
    help="project-id",
    type=int,
    required=True,
)
@click.option(
    "-p",
    "--page",
    help="page",
    type=int,
    default=1,
)
def list_chats(project_id, page):
    """list-chats"""
    out = m_main_get_conversation_record_list(project_id=project_id, page=page)
    click.echo(out)

@click.command(name="sub-list-chats")
@click.option(
    "-project-id",
    "--project-id",
    help="project-id",
    type=int,
    required=True,
)
@click.option(
    "-sub-user-id",
    "--sub-user-id",
    help="sub-user-id",
    type=int,
    required=True,
)
@click.option(
    "-p",
    "--page",
    help="page",
    type=int,
    default=1,
)
def sub_list_chats(project_id, sub_user_id, page):
    """sub-list-chats"""
    out = m_sub_get_conversation_record_list(sub_user_id, project_id, page)
    click.echo(out)
@click.command(name="view-chat")
def view_chat():
    """view-chat"""
    pass
@click.command(name="reply-chat")
def reply_chat():
    """reply-chat"""
    pass

# Device
@click.command(name="register-device")
def register_device():
    """register-device"""
    pass
@click.command(name="fetch-device-task")
def fetch_device_task():
    """fetch-device-task"""
    pass
@click.command(name="report-task-result")
def report_task_result():
    """report-task-result"""
    pass
@click.command(name="report-receive-content")
def report_receive_content():
    """report-receive-content"""
    pass

sms_cli.add_command(list_project)
sms_cli.add_command(create_project)
sms_cli.add_command(delete_project)
sms_cli.add_command(update_project)
sms_cli.add_command(allocate_device_to_project)
sms_cli.add_command(create_task)
sms_cli.add_command(delete_task)
sms_cli.add_command(list_tasks)
sms_cli.add_command(sub_list_tasks)
sms_cli.add_command(list_task_record)
sms_cli.add_command(list_groups)
sms_cli.add_command(create_group)
sms_cli.add_command(allocate_device_to_group)
sms_cli.add_command(update_group)
sms_cli.add_command(delete_group)
sms_cli.add_command(list_chats)
sms_cli.add_command(view_chat)
sms_cli.add_command(reply_chat)
sms_cli.add_command(register_device)
sms_cli.add_command(fetch_device_task)
sms_cli.add_command(report_task_result)
sms_cli.add_command(report_receive_content)
sms_cli.add_command(list_device)
sms_cli.add_command(sub_list_device)
sms_cli.add_command(download_task)
sms_cli.add_command(sub_list_task_record)
sms_cli.add_command(sub_list_chats)

if __name__ == '__main__':
    sms_cli()