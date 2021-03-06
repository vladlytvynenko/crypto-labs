# Generated by Django 3.1.4 on 2020-12-08 20:00

from django.db import migrations
import encrypted_fields.fields


class Migration(migrations.Migration):

    dependencies = [
        ('accounts', '0002_auto_20201208_1951'),
    ]

    operations = [
        migrations.AlterField(
            model_name='user',
            name='address',
            field=encrypted_fields.fields.EncryptedCharField(blank=True, help_text='Your address', max_length=200, null=True),
        ),
        migrations.AlterField(
            model_name='user',
            name='phone',
            field=encrypted_fields.fields.EncryptedCharField(blank=True, help_text='Your phone number', max_length=48, null=True),
        ),
    ]
