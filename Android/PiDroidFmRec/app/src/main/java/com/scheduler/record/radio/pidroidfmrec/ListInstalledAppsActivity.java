package com.scheduler.record.radio.pidroidfmrec;

import android.content.Intent;
import android.content.pm.PackageInfo;
import android.content.pm.PackageManager;
import android.os.Bundle;
import android.support.annotation.NonNull;
import android.support.v7.app.AppCompatActivity;
import android.view.Menu;
import android.view.MenuInflater;
import android.view.MenuItem;
import android.view.Window;
import android.view.WindowManager;
import android.widget.TableLayout;
import android.widget.TableRow;
import android.widget.TextView;

import java.util.List;

public class ListInstalledAppsActivity extends AppCompatActivity
{

    private static final String TAG = "ListAppsActivity";

    @Override
    protected void onCreate(Bundle savedInstanceState)
    {
        super.onCreate(savedInstanceState);
        // Remove title bar.
        this.requestWindowFeature(Window.FEATURE_NO_TITLE);
        // Remove notification bar.
        this.getWindow().setFlags(WindowManager.LayoutParams.FLAG_FULLSCREEN, WindowManager.LayoutParams.FLAG_FULLSCREEN);
        setContentView(R.layout.activity_list_installed_apps);

        TableLayout tableLayout = (TableLayout) findViewById(R.id.tableLayout_table);

        final PackageManager pm = getPackageManager();
        // Get a list of installed apps.
        List<PackageInfo> packages = pm.getInstalledPackages(PackageManager.GET_PERMISSIONS);

        assert tableLayout != null;
        /*
        For each app, get the package name and its corresponding app name add it to the
        table view.
        */
        for (PackageInfo packageInfo : packages)
        {
            TableRow tableRow = new TableRow(this);

            TextView textViewAppName = getTextView(pm.getApplicationLabel(packageInfo.applicationInfo).toString());
            TextView textViewPackageName = getTextView(packageInfo.packageName);

            tableRow.addView(textViewAppName);
            tableRow.addView(textViewPackageName);

            tableLayout.addView(tableRow);
        }
    }

    @Override
    public boolean onCreateOptionsMenu(Menu menu)
    {
        MenuInflater menuInflater = getMenuInflater();
        menuInflater.inflate(R.menu.options_menu, menu);
        return true;
    }

    @Override
    public boolean onOptionsItemSelected(MenuItem item)
    {
        switch (item.getItemId())
        {
            case R.id.set_wake_up:
                startActivity(new Intent(ListInstalledAppsActivity.this, MainActivity.class));
                return true;
            case R.id.show_device_ip:
                ShowIpDialog ipDialog = new ShowIpDialog();
                ipDialog.show(getFragmentManager(), TAG);
                return true;
        }
        return true;
    }

    @NonNull
    private TextView getTextView(String text)
    {
        TextView textViewAppName = new TextView(this);
        textViewAppName.setText(text);
        return textViewAppName;
    }
}
