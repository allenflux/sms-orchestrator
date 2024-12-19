import click


@click.group()
def sms_cli():
    """sms_cli"""
    pass

# Project
@click.command(name="list-project")
def list_project():
    """list-project"""
    pass
@click.command(name="create-project")
def create_project():
    """create-project"""
    pass
@click.command(name="delete-project")
def delete_project():
    """delete-project"""
    pass
@click.command(name="update-projec")
def update_project():
    """update-project"""
    pass
@click.command(name="allocate-device-to-project")
def allocate_device_to_project():
    """allocate-device-to-project"""
    pass

# Task
@click.command(name="create-task")
def create_task():
    """create-task"""
    pass
@click.command(name="delete-task")
def delete_task():
    """delete-task"""
    pass
@click.command(name="list-tasks")
def list_tasks():
    """list-tasks"""
    pass
@click.command(name="list-task-record")
def list_task_record():
    """list-task-record"""
    pass

# Group
@click.command(name="list-groups")
def list_groups():
    """list-groups"""
    pass
@click.command(name="create-group")
def create_group():
    """create-group"""
    pass
@click.command(name="allocate-device-to-project")
def allocate_device_to_group():
    """allocate-device-to-project"""
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
def list_chats():
    """list-chats"""
    pass
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

if __name__ == '__main__':
    sms_cli()