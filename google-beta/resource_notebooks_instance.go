// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceNotebooksInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceNotebooksInstanceCreate,
		Read:   resourceNotebooksInstanceRead,
		Update: resourceNotebooksInstanceUpdate,
		Delete: resourceNotebooksInstanceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceNotebooksInstanceImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(15 * time.Minute),
			Update: schema.DefaultTimeout(15 * time.Minute),
			Delete: schema.DefaultTimeout(15 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"location": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to the zone where the machine resides.`,
			},
			"machine_type": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `A reference to a machine type which defines VM kind.`,
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The name specified for the Notebook instance.`,
			},
			"accelerator_config": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The hardware accelerator used on this instance. If you use accelerators, 
make sure that your configuration has enough vCPUs and memory to support the 
machineType you have selected.`,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"core_count": {
							Type:        schema.TypeInt,
							Required:    true,
							ForceNew:    true,
							Description: `Count of cores of this accelerator.`,
						},
						"type": {
							Type:         schema.TypeString,
							Required:     true,
							ForceNew:     true,
							ValidateFunc: validation.StringInSlice([]string{"ACCELERATOR_TYPE_UNSPECIFIED", "NVIDIA_TESLA_K80", "NVIDIA_TESLA_P100", "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_T4_VWS", "NVIDIA_TESLA_P100_VWS", "NVIDIA_TESLA_P4_VWS", "TPU_V2", "TPU_V3"}, false),
							Description:  `Type of this accelerator. Possible values: ["ACCELERATOR_TYPE_UNSPECIFIED", "NVIDIA_TESLA_K80", "NVIDIA_TESLA_P100", "NVIDIA_TESLA_V100", "NVIDIA_TESLA_P4", "NVIDIA_TESLA_T4", "NVIDIA_TESLA_T4_VWS", "NVIDIA_TESLA_P100_VWS", "NVIDIA_TESLA_P4_VWS", "TPU_V2", "TPU_V3"]`,
						},
					},
				},
			},
			"boot_disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The size of the boot disk in GB attached to this instance, 
up to a maximum of 64000 GB (64 TB). The minimum recommended value is 100 GB. 
If not specified, this defaults to 100.`,
			},
			"boot_disk_type": {
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validation.StringInSlice([]string{"DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", ""}, false),
				Description:  `Possible disk types for notebook instances. Possible values: ["DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED"]`,
			},
			"container_image": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Use a container image to start the notebook instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"repository": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The path to the container image repository. 
For example: gcr.io/{project_id}/{imageName}`,
						},
						"tag": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The tag of the container image. If not specified, this defaults to the latest tag.`,
						},
					},
				},
				ExactlyOneOf: []string{"vm_image", "container_image"},
			},
			"custom_gpu_driver_path": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Specify a custom Cloud Storage path where the GPU driver is stored. 
If not specified, we'll automatically choose from official GPU drivers.`,
			},
			"data_disk_size_gb": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Description: `The size of the data disk in GB attached to this instance, 
up to a maximum of 64000 GB (64 TB). 
You can choose the size of the data disk based on how big your notebooks and data are. 
If not specified, this defaults to 100.`,
			},
			"data_disk_type": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("DISK_TYPE_UNSPECIFIED"),
				Description:      `Possible disk types for notebook instances. Possible values: ["DISK_TYPE_UNSPECIFIED", "PD_STANDARD", "PD_SSD", "PD_BALANCED"]`,
			},
			"disk_encryption": {
				Type:             schema.TypeString,
				Optional:         true,
				ForceNew:         true,
				ValidateFunc:     validation.StringInSlice([]string{"DISK_ENCRYPTION_UNSPECIFIED", "GMEK", "CMEK", ""}, false),
				DiffSuppressFunc: emptyOrDefaultStringSuppress("DISK_ENCRYPTION_UNSPECIFIED"),
				Description:      `Disk encryption method used on the boot and data disks, defaults to GMEK. Possible values: ["DISK_ENCRYPTION_UNSPECIFIED", "GMEK", "CMEK"]`,
			},
			"install_gpu_driver": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
				Description: `Whether the end user authorizes Google Cloud to install GPU driver
on this instance. If this field is empty or set to false, the GPU driver
won't be installed. Only applicable to instances with GPUs.`,
			},
			"instance_owners": {
				Type:     schema.TypeList,
				Optional: true,
				ForceNew: true,
				Description: `The list of owners of this instance after creation. 
Format: alias@example.com.
Currently supports one owner only. 
If not specified, all of the service account users of 
your VM instance's service account can use the instance.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"kms_key": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `The KMS key used to encrypt the disks, only applicable if diskEncryption is CMEK. 
Format: projects/{project_id}/locations/{location}/keyRings/{key_ring_id}/cryptoKeys/{key_id}`,
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Labels to apply to this instance. These can be later modified by the setLabels method.
An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"metadata": {
				Type:     schema.TypeMap,
				Optional: true,
				ForceNew: true,
				Description: `Custom metadata to apply to this instance.
An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"network": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The name of the VPC that this instance is in. 
Format: projects/{project_id}/global/networks/{network_id}`,
			},
			"no_proxy_access": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `the notebook instance will not register with the proxy..`,
			},
			"no_public_ip": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `no public IP will be assigned to this instance.`,
			},
			"no_remove_data_disk": {
				Type:        schema.TypeBool,
				Optional:    true,
				ForceNew:    true,
				Description: `If true, the data disk will not be auto deleted when deleting the instance.`,
			},
			"post_startup_script": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Description: `Path to a Bash script that automatically runs after a 
notebook instance fully boots up. The path must be a URL 
or Cloud Storage path (gs://path-to-file/file-name).`,
			},
			"service_account": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
				ForceNew: true,
				Description: `The service account on this instance, giving access to other 
Google Cloud services. You can use any service account within 
the same project, but you must have the service account user 
permission to use the instance. If not specified, 
the Compute Engine default service account is used.`,
			},
			"subnet": {
				Type:             schema.TypeString,
				Computed:         true,
				Optional:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description: `The name of the subnet that this instance is in. 
Format: projects/{project_id}/regions/{region}/subnetworks/{subnetwork_id}`,
			},
			"vm_image": {
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: `Use a Compute Engine VM image to start the notebook instance.`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							Description: `The name of the Google Cloud project that this VM image belongs to. 
Format: projects/{project_id}`,
						},
						"image_family": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Use this VM image family to find the image; the newest image in this family will be used.`,
						},
						"image_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `Use VM image name to find the image.`,
						},
					},
				},
				ExactlyOneOf: []string{"vm_image", "container_image"},
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `Instance creation time`,
			},
			"proxy_uri": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The proxy endpoint that is used to access the Jupyter notebook.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The state of this instance.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Optional:    true,
				Description: `Instance update time.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
	}
}

func resourceNotebooksInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	obj := make(map[string]interface{})
	machineTypeProp, err := expandNotebooksInstanceMachineType(d.Get("machine_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("machine_type"); !isEmptyValue(reflect.ValueOf(machineTypeProp)) && (ok || !reflect.DeepEqual(v, machineTypeProp)) {
		obj["machineType"] = machineTypeProp
	}
	postStartupScriptProp, err := expandNotebooksInstancePostStartupScript(d.Get("post_startup_script"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("post_startup_script"); !isEmptyValue(reflect.ValueOf(postStartupScriptProp)) && (ok || !reflect.DeepEqual(v, postStartupScriptProp)) {
		obj["postStartupScript"] = postStartupScriptProp
	}
	instanceOwnersProp, err := expandNotebooksInstanceInstanceOwners(d.Get("instance_owners"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("instance_owners"); !isEmptyValue(reflect.ValueOf(instanceOwnersProp)) && (ok || !reflect.DeepEqual(v, instanceOwnersProp)) {
		obj["instanceOwners"] = instanceOwnersProp
	}
	serviceAccountProp, err := expandNotebooksInstanceServiceAccount(d.Get("service_account"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("service_account"); !isEmptyValue(reflect.ValueOf(serviceAccountProp)) && (ok || !reflect.DeepEqual(v, serviceAccountProp)) {
		obj["serviceAccount"] = serviceAccountProp
	}
	acceleratorConfigProp, err := expandNotebooksInstanceAcceleratorConfig(d.Get("accelerator_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("accelerator_config"); !isEmptyValue(reflect.ValueOf(acceleratorConfigProp)) && (ok || !reflect.DeepEqual(v, acceleratorConfigProp)) {
		obj["acceleratorConfig"] = acceleratorConfigProp
	}
	installGpuDriverProp, err := expandNotebooksInstanceInstallGpuDriver(d.Get("install_gpu_driver"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("install_gpu_driver"); !isEmptyValue(reflect.ValueOf(installGpuDriverProp)) && (ok || !reflect.DeepEqual(v, installGpuDriverProp)) {
		obj["installGpuDriver"] = installGpuDriverProp
	}
	customGpuDriverPathProp, err := expandNotebooksInstanceCustomGpuDriverPath(d.Get("custom_gpu_driver_path"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_gpu_driver_path"); !isEmptyValue(reflect.ValueOf(customGpuDriverPathProp)) && (ok || !reflect.DeepEqual(v, customGpuDriverPathProp)) {
		obj["customGpuDriverPath"] = customGpuDriverPathProp
	}
	bootDiskTypeProp, err := expandNotebooksInstanceBootDiskType(d.Get("boot_disk_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("boot_disk_type"); !isEmptyValue(reflect.ValueOf(bootDiskTypeProp)) && (ok || !reflect.DeepEqual(v, bootDiskTypeProp)) {
		obj["bootDiskType"] = bootDiskTypeProp
	}
	bootDiskSizeGbProp, err := expandNotebooksInstanceBootDiskSizeGb(d.Get("boot_disk_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("boot_disk_size_gb"); !isEmptyValue(reflect.ValueOf(bootDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, bootDiskSizeGbProp)) {
		obj["bootDiskSizeGb"] = bootDiskSizeGbProp
	}
	dataDiskTypeProp, err := expandNotebooksInstanceDataDiskType(d.Get("data_disk_type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_disk_type"); !isEmptyValue(reflect.ValueOf(dataDiskTypeProp)) && (ok || !reflect.DeepEqual(v, dataDiskTypeProp)) {
		obj["dataDiskType"] = dataDiskTypeProp
	}
	dataDiskSizeGbProp, err := expandNotebooksInstanceDataDiskSizeGb(d.Get("data_disk_size_gb"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("data_disk_size_gb"); !isEmptyValue(reflect.ValueOf(dataDiskSizeGbProp)) && (ok || !reflect.DeepEqual(v, dataDiskSizeGbProp)) {
		obj["dataDiskSizeGb"] = dataDiskSizeGbProp
	}
	noRemoveDataDiskProp, err := expandNotebooksInstanceNoRemoveDataDisk(d.Get("no_remove_data_disk"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("no_remove_data_disk"); !isEmptyValue(reflect.ValueOf(noRemoveDataDiskProp)) && (ok || !reflect.DeepEqual(v, noRemoveDataDiskProp)) {
		obj["noRemoveDataDisk"] = noRemoveDataDiskProp
	}
	diskEncryptionProp, err := expandNotebooksInstanceDiskEncryption(d.Get("disk_encryption"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("disk_encryption"); !isEmptyValue(reflect.ValueOf(diskEncryptionProp)) && (ok || !reflect.DeepEqual(v, diskEncryptionProp)) {
		obj["diskEncryption"] = diskEncryptionProp
	}
	kmsKeyProp, err := expandNotebooksInstanceKmsKey(d.Get("kms_key"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("kms_key"); !isEmptyValue(reflect.ValueOf(kmsKeyProp)) && (ok || !reflect.DeepEqual(v, kmsKeyProp)) {
		obj["kmsKey"] = kmsKeyProp
	}
	noPublicIpProp, err := expandNotebooksInstanceNoPublicIp(d.Get("no_public_ip"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("no_public_ip"); !isEmptyValue(reflect.ValueOf(noPublicIpProp)) && (ok || !reflect.DeepEqual(v, noPublicIpProp)) {
		obj["noPublicIp"] = noPublicIpProp
	}
	noProxyAccessProp, err := expandNotebooksInstanceNoProxyAccess(d.Get("no_proxy_access"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("no_proxy_access"); !isEmptyValue(reflect.ValueOf(noProxyAccessProp)) && (ok || !reflect.DeepEqual(v, noProxyAccessProp)) {
		obj["noProxyAccess"] = noProxyAccessProp
	}
	networkProp, err := expandNotebooksInstanceNetwork(d.Get("network"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("network"); !isEmptyValue(reflect.ValueOf(networkProp)) && (ok || !reflect.DeepEqual(v, networkProp)) {
		obj["network"] = networkProp
	}
	subnetProp, err := expandNotebooksInstanceSubnet(d.Get("subnet"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("subnet"); !isEmptyValue(reflect.ValueOf(subnetProp)) && (ok || !reflect.DeepEqual(v, subnetProp)) {
		obj["subnet"] = subnetProp
	}
	labelsProp, err := expandNotebooksInstanceLabels(d.Get("labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	metadataProp, err := expandNotebooksInstanceMetadata(d.Get("metadata"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("metadata"); !isEmptyValue(reflect.ValueOf(metadataProp)) && (ok || !reflect.DeepEqual(v, metadataProp)) {
		obj["metadata"] = metadataProp
	}
	vmImageProp, err := expandNotebooksInstanceVmImage(d.Get("vm_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("vm_image"); !isEmptyValue(reflect.ValueOf(vmImageProp)) && (ok || !reflect.DeepEqual(v, vmImageProp)) {
		obj["vmImage"] = vmImageProp
	}
	containerImageProp, err := expandNotebooksInstanceContainerImage(d.Get("container_image"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("container_image"); !isEmptyValue(reflect.ValueOf(containerImageProp)) && (ok || !reflect.DeepEqual(v, containerImageProp)) {
		obj["containerImage"] = containerImageProp
	}

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances?instanceId={{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Instance: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating Instance: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = notebooksOperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating Instance",
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Instance: %s", err)
	}

	// This may have caused the ID to update - update it if so.
	id, err = replaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating Instance %q: %#v", d.Id(), res)

	return resourceNotebooksInstanceRead(d, meta)
}

func resourceNotebooksInstanceRead(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("NotebooksInstance %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	if err := d.Set("machine_type", flattenNotebooksInstanceMachineType(res["machineType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("post_startup_script", flattenNotebooksInstancePostStartupScript(res["postStartupScript"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("proxy_uri", flattenNotebooksInstanceProxyUri(res["proxyUri"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("service_account", flattenNotebooksInstanceServiceAccount(res["serviceAccount"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("accelerator_config", flattenNotebooksInstanceAcceleratorConfig(res["acceleratorConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("state", flattenNotebooksInstanceState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("install_gpu_driver", flattenNotebooksInstanceInstallGpuDriver(res["installGpuDriver"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("custom_gpu_driver_path", flattenNotebooksInstanceCustomGpuDriverPath(res["customGpuDriverPath"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("data_disk_type", flattenNotebooksInstanceDataDiskType(res["dataDiskType"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("no_remove_data_disk", flattenNotebooksInstanceNoRemoveDataDisk(res["noRemoveDataDisk"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("disk_encryption", flattenNotebooksInstanceDiskEncryption(res["diskEncryption"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("kms_key", flattenNotebooksInstanceKmsKey(res["kmsKey"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("no_public_ip", flattenNotebooksInstanceNoPublicIp(res["noPublicIp"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("no_proxy_access", flattenNotebooksInstanceNoProxyAccess(res["noProxyAccess"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("network", flattenNotebooksInstanceNetwork(res["network"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("subnet", flattenNotebooksInstanceSubnet(res["subnet"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("labels", flattenNotebooksInstanceLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("create_time", flattenNotebooksInstanceCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}
	if err := d.Set("update_time", flattenNotebooksInstanceUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Instance: %s", err)
	}

	return nil
}

func resourceNotebooksInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	d.Partial(true)

	if d.HasChange("labels") {
		obj := make(map[string]interface{})

		labelsProp, err := expandNotebooksInstanceLabels(d.Get("labels"), d, config)
		if err != nil {
			return err
		} else if v, ok := d.GetOkExists("labels"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
			obj["labels"] = labelsProp
		}

		url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances/{{name}}:setLabels")
		if err != nil {
			return err
		}

		// err == nil indicates that the billing_project value was found
		if bp, err := getBillingProject(d, config); err == nil {
			billingProject = bp
		}

		res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, obj, d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return fmt.Errorf("Error updating Instance %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Instance %q: %#v", d.Id(), res)
		}

		err = notebooksOperationWaitTime(
			config, res, project, "Updating Instance",
			d.Timeout(schema.TimeoutUpdate))
		if err != nil {
			return err
		}
	}

	d.Partial(false)

	return resourceNotebooksInstanceRead(d, meta)
}

func resourceNotebooksInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	var m providerMeta

	err := d.GetProviderMeta(&m)
	if err != nil {
		return err
	}

	config := meta.(*Config)
	config.userAgent = fmt.Sprintf("%s %s", config.userAgent, m.ModuleName)

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{NotebooksBasePath}}projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting Instance %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "Instance")
	}

	err = notebooksOperationWaitTime(
		config, res, project, "Deleting Instance",
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Instance %q: %#v", d.Id(), res)
	return nil
}

func resourceNotebooksInstanceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/instances/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<name>[^/]+)",
		"(?P<location>[^/]+)/(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/locations/{{location}}/instances/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenNotebooksInstanceMachineType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return NameFromSelfLinkStateFunc(v)
}

func flattenNotebooksInstancePostStartupScript(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceProxyUri(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceServiceAccount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceAcceleratorConfig(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["type"] =
		flattenNotebooksInstanceAcceleratorConfigType(original["type"], d, config)
	transformed["core_count"] =
		flattenNotebooksInstanceAcceleratorConfigCoreCount(original["coreCount"], d, config)
	return []interface{}{transformed}
}
func flattenNotebooksInstanceAcceleratorConfigType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceAcceleratorConfigCoreCount(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := strconv.ParseInt(strVal, 10, 64); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenNotebooksInstanceState(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceInstallGpuDriver(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceCustomGpuDriverPath(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceDataDiskType(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceNoRemoveDataDisk(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceDiskEncryption(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceKmsKey(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceNoPublicIp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceNoProxyAccess(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceNetwork(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceSubnet(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceLabels(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceCreateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenNotebooksInstanceUpdateTime(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func expandNotebooksInstanceMachineType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstancePostStartupScript(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstanceOwners(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceServiceAccount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfig(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedType, err := expandNotebooksInstanceAcceleratorConfigType(original["type"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedType); val.IsValid() && !isEmptyValue(val) {
		transformed["type"] = transformedType
	}

	transformedCoreCount, err := expandNotebooksInstanceAcceleratorConfigCoreCount(original["core_count"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedCoreCount); val.IsValid() && !isEmptyValue(val) {
		transformed["coreCount"] = transformedCoreCount
	}

	return transformed, nil
}

func expandNotebooksInstanceAcceleratorConfigType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceAcceleratorConfigCoreCount(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceInstallGpuDriver(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceCustomGpuDriverPath(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceBootDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskType(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDataDiskSizeGb(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoRemoveDataDisk(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceDiskEncryption(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceKmsKey(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoPublicIp(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNoProxyAccess(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceNetwork(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceSubnet(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceLabels(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNotebooksInstanceMetadata(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandNotebooksInstanceVmImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedProject, err := expandNotebooksInstanceVmImageProject(original["project"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedProject); val.IsValid() && !isEmptyValue(val) {
		transformed["project"] = transformedProject
	}

	transformedImageFamily, err := expandNotebooksInstanceVmImageImageFamily(original["image_family"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageFamily); val.IsValid() && !isEmptyValue(val) {
		transformed["imageFamily"] = transformedImageFamily
	}

	transformedImageName, err := expandNotebooksInstanceVmImageImageName(original["image_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedImageName); val.IsValid() && !isEmptyValue(val) {
		transformed["imageName"] = transformedImageName
	}

	return transformed, nil
}

func expandNotebooksInstanceVmImageProject(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageFamily(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceVmImageImageName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImage(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRepository, err := expandNotebooksInstanceContainerImageRepository(original["repository"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRepository); val.IsValid() && !isEmptyValue(val) {
		transformed["repository"] = transformedRepository
	}

	transformedTag, err := expandNotebooksInstanceContainerImageTag(original["tag"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedTag); val.IsValid() && !isEmptyValue(val) {
		transformed["tag"] = transformedTag
	}

	return transformed, nil
}

func expandNotebooksInstanceContainerImageRepository(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandNotebooksInstanceContainerImageTag(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}
