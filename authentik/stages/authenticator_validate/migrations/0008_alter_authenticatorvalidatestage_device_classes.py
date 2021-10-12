# Generated by Django 3.2.3 on 2021-05-29 16:10

import django.contrib.postgres.fields
from django.db import migrations, models

import authentik.stages.authenticator_validate.models


class Migration(migrations.Migration):

    dependencies = [
        ("authentik_stages_authenticator_validate", "0007_auto_20210403_0927"),
    ]

    operations = [
        migrations.AlterField(
            model_name="authenticatorvalidatestage",
            name="device_classes",
            field=django.contrib.postgres.fields.ArrayField(
                base_field=models.TextField(
                    choices=[
                        ("static", "Static"),
                        ("totp", "TOTP"),
                        ("webauthn", "WebAuthn"),
                        ("duo", "Duo"),
                        ("sms", "SMS"),
                    ]
                ),
                default=authentik.stages.authenticator_validate.models.default_device_classes,
                help_text="Device classes which can be used to authenticate",
                size=None,
            ),
        ),
    ]
