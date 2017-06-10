#include "searchwindow.h"
#include "ui_searchwindow.h"

SearchWindow::SearchWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::SearchWindow)
{
    ui->setupUi(this);
}

SearchWindow::~SearchWindow()
{
    delete ui;
}

void SearchWindow::on_snsearchB_clicked()
{
    QString data = ui->snF->text();
    emit plsSearchStudName(data);
}

void SearchWindow::on_sssearchB_clicked()
{
    QString data = ui->srF->text();
    emit plsSearchStudSur(data);
}

void SearchWindow::on_cnsearchB_clicked()
{
    QString data = ui->crF->text();
    emit plsSearchCours(data);
}
